package adapters

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/adjust/rmq/v4"
)

var queues = map[string]rmq.Queue{}

type MessageQueue struct {
	Connection rmq.Connection
}

type ConsumerConfig struct {
	PrefetchCount int
	PollInterval time.Duration
	QueueName string
	Callback rmq.ConsumerFunc
}

func (q MessageQueue) getQueue(queueName string) (rmq.Queue, error) {
	if queues[queueName] == nil {
		queue, err := q.Connection.OpenQueue(queueName)
		if err != nil {
			return nil, err
		}
		queues[queueName] = queue
		return queue, nil
	}
	return queues[queueName], nil
}

func (q MessageQueue) Publish(queueName string, msg interface{}) error {
	queue, err := q.getQueue(queueName)
	if err != nil {
		return err
	}

	fmt.Printf("publish msg to queue %v\n", queueName)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return queue.PublishBytes(msgBytes)
}

func (q MessageQueue) Consume(consumerConfig ConsumerConfig) error {
	queue, err := q.getQueue(consumerConfig.QueueName)
	if err != nil {
		return err
	}

	err = queue.StartConsuming(int64(consumerConfig.PrefetchCount), consumerConfig.PollInterval)
	if err != nil {
		return err
	}

	consumerName, err := queue.AddConsumerFunc(consumerConfig.QueueName, consumerConfig.Callback)
	if err != nil {
		return err
	}

	fmt.Printf("start consume queue %v with consumer %v\n", consumerConfig.QueueName, consumerName)
	return nil
}

func InitializeMessageQueue(redisUri string) *MessageQueue {
	connection, err := rmq.OpenConnection("edwork-api-go", "tcp", redisUri, 1, nil)
	if err != nil {
		fmt.Println("error creating message queue: ", err)
		os.Exit(1)
	}

	return &MessageQueue{
		Connection: connection,
	}
}
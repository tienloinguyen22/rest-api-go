package consumers

import (
	"fmt"

	"github.com/adjust/rmq/v4"
)

type ConsumerService struct {}

func NewConsumerService() *ConsumerService {
	return &ConsumerService{}
}

func (s ConsumerService) ResizeImage(delivery rmq.Delivery) {
	fmt.Println("Resize image!!")
	delivery.Ack()
}
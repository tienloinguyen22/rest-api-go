package consumers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/adjust/rmq/v4"
	"github.com/h2non/bimg"
	"github.com/tienloinguyen22/edwork-api-go/utils"
)

type ConsumerService struct {}

func NewConsumerService() *ConsumerService {
	return &ConsumerService{}
}

func (s ConsumerService) ResizeImage(delivery rmq.Delivery) {
	var payload ResizeImagePayload
	if err := json.Unmarshal([]byte(delivery.Payload()), &payload); err != nil {
		fmt.Println("error parse msg payload: ", err)
		delivery.Ack()
		return
	}

	buffer, err := bimg.Read("./temp/" + payload.Filename)
	if err != nil {
		fmt.Println("error reading image file: ", err)
		delivery.Ack()
		return
	}

	newImage, err := bimg.NewImage(buffer).Process(bimg.Options{
		Lossless: true,
		Quality: 60,
	})
	if err != nil {
		fmt.Println("error processing image file: ", err)
		delivery.Ack()
		return
	}

	err = utils.EnsureFolderExist("./uploads")
	if err != nil {
		fmt.Println("error finding /uploads folder: ", err)
		delivery.Ack()
		return
	}

	err = bimg.Write("./uploads/" + payload.Filename, newImage)
	if err != nil {
		fmt.Println("error writing image file: ", err)
		delivery.Ack()
		return
	}

	err = os.Remove("./temp/" + payload.Filename)
	if err != nil {
		fmt.Println("error remove image file from /temp: ", err)
		delivery.Ack()
		return
	}

	delivery.Ack()
	fmt.Println("finish resize image: " + payload.Filename)
}
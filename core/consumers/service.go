package consumers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"github.com/adjust/rmq/v4"
	"github.com/h2non/bimg"
	"github.com/tienloinguyen22/edwork-api-go/adapters"
	"github.com/tienloinguyen22/edwork-api-go/utils"
)

type ConsumerService struct {
	EmailClient *adapters.EmailClient
}

func NewConsumerService(emailClient *adapters.EmailClient) *ConsumerService {
	return &ConsumerService{
		EmailClient: emailClient,
	}
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

func (s ConsumerService) SendMail(delivery rmq.Delivery) {
	var payload SendMailPayload
	if err := json.Unmarshal([]byte(delivery.Payload()), &payload); err != nil {
		fmt.Println("error parse msg payload: ", err)
		delivery.Ack()
		return
	}

	tmpl, err := template.ParseFiles("./templates/" + payload.Template)
	if err != nil {
		fmt.Println("error parsing html template: ", err)
		delivery.Ack()
		return
	}

	var content bytes.Buffer
	if err := tmpl.Execute(&content, payload.MailData); err != nil {
		fmt.Println("error execute html template: ", err)
		delivery.Ack()
		return
	}

	subject := ""
	if payload.Template == "forgot-password.html" {
		subject = "[EdWORK] Thay đổi mật khẩu"
	}

	s.EmailClient.SendMail(&adapters.SendMailPayload{
		From: "Neoflies <tienloinguyen22@gmail.com>",
		To: payload.To,
		CC: payload.CC,
		Subject: subject,
		Body: content.String(),
		Attachements: payload.Attachements,
	})
}
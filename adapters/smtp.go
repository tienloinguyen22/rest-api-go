package adapters

import (
	"fmt"
	"os"

	mail "github.com/xhit/go-simple-mail/v2"
)

type EmailClient struct {
	SmtpClient *mail.SMTPClient
}

type SendMailPayload struct {
	From string
	To []string
	CC []string
	Subject string
	Body string
	Attachements []*mail.File
}

func (s EmailClient) SendMail(payload *SendMailPayload) error {
	email := mail.NewMSG()
	email.SetFrom(payload.From)
	email.SetSubject(payload.Subject)
	email.SetBody(mail.TextHTML, payload.Body)
	for _, to := range payload.To {
		email.AddTo(to)
	}
	for _, cc := range payload.CC {
		email.AddCc(cc)
	}
	for _, attachment := range payload.Attachements {
		email.Attach(attachment)
	}

	return email.Send(s.SmtpClient)
}

func InitializeSmtpClient(host string, port int, username string, password string) *EmailClient {
	server := mail.NewSMTPClient()
	server.Host = host
	server.Port = port
	server.Username = username
	server.Password = password
	server.Encryption = mail.EncryptionSSL

	smtpClient, err := server.Connect()
	if err != nil {
		fmt.Println("error connecting to smtp servier: ", err)
		os.Exit(1)
	}

	return &EmailClient{
		SmtpClient: smtpClient,
	}
}
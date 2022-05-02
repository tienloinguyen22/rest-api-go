package consumers

import mail "github.com/xhit/go-simple-mail/v2"

type ResizeImagePayload struct {
	Filename string
}

type SendMailPayload struct {
	Template string
	MailData interface{}
	To []string
	CC []string
	Attachements []*mail.File
}
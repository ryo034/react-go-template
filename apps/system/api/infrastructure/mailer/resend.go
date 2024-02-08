package mailer

import (
	"github.com/resend/resend-go/v2"
)

type ResendMailer struct {
	client *resend.Client
}

func NewResendMailer(apiKey string) Client {
	c := resend.NewClient(apiKey)
	return &ResendMailer{c}
}

func (m *ResendMailer) Send(conf Config) error {
	params := &resend.SendEmailRequest{
		From:    conf.From,
		To:      conf.To,
		Html:    conf.Body,
		Subject: conf.Subject,
		Cc:      conf.Cc,
		Bcc:     conf.Bcc,
		ReplyTo: conf.ReplyTo,
	}
	_, err := m.client.Emails.Send(params)
	if err != nil {
		return err
	}
	return nil
}

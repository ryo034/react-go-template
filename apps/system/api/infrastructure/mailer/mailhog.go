package mailer

import (
	"fmt"
	"net/smtp"
	"strings"
)

type MailhogMailer struct {
	Host string
	Port int
}

func NewMailhogMailer(host string, port int) Client {
	return &MailhogMailer{host, port}
}

func (m *MailhogMailer) Send(conf Config) error {
	addr := fmt.Sprintf("%s:%d", m.Host, m.Port)

	header := make(map[string]string)
	header["From"] = conf.From
	header["To"] = strings.Join(conf.To, ", ")
	header["Subject"] = conf.Subject
	if len(conf.Cc) > 0 {
		header["Cc"] = strings.Join(conf.Cc, ", ")
	}
	if conf.ReplyTo != "" {
		header["Reply-To"] = conf.ReplyTo
	}

	if conf.IsHTML {
		header["Content-Type"] = "text/html; charset=UTF-8"
	} else {
		header["Content-Type"] = "text/plain; charset=UTF-8"
	}

	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + conf.Body

	return smtp.SendMail(addr, nil, conf.From, append(conf.To, append(conf.Cc, conf.Bcc...)...), []byte(msg))
}

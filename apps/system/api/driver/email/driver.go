package email

import (
	"bytes"
	"context"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/mailer"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"golang.org/x/text/language"
	"html/template"
)

type Driver interface {
	SendOTP(ctx context.Context, mailTo account.Email, code string) error
}

type driver struct {
	co           shared.ContextOperator
	mc           mailer.Client
	noreplyEmail account.Email
}

func NewDriver(co shared.ContextOperator, mc mailer.Client, noreplyEmail account.Email) Driver {
	return &driver{co, mc, noreplyEmail}
}

type otpTemplateData struct {
	Code string
}

func (d *driver) SendOTP(ctx context.Context, mailTo account.Email, code string) error {
	lang, err := d.co.GetLang(ctx)
	if err != nil {
		return err
	}
	tmplPath := ""
	subject := ""
	switch lang {
	case language.English:
		subject = "Your OTP Code"
		tmplPath = "driver/email/otp_template_en.html"
		if err != nil {
			return err
		}
	default:
		subject = "あなたのOTPコード"
		tmplPath = "driver/email/otp_template_ja.html"
		if err != nil {
			return err
		}
	}
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}
	data := otpTemplateData{Code: code}
	var body bytes.Buffer
	if err = tmpl.Execute(&body, data); err != nil {
		return err
	}

	conf := mailer.Config{
		From:    d.noreplyEmail.ToString(),
		To:      []string{mailTo.ToString()},
		Subject: subject,
		Body:    body.String(),
		Cc:      nil,
		Bcc:     nil,
		ReplyTo: "",
		IsHTML:  true,
	}
	return d.mc.Send(conf)
}

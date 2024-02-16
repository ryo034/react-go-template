package email

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace"
	"github.com/ryo034/react-go-template/apps/system/api/domain/workspace/invitation"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/mailer"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"golang.org/x/text/language"
	"html/template"
)

type Driver interface {
	SendOTP(ctx context.Context, mailTo account.Email, code string) error
	SendInvite(ctx context.Context, inviter workspace.Inviter, i *invitation.Invitation) error
}

type driver struct {
	serviceName  string
	co           shared.ContextOperator
	mc           mailer.Client
	noReplyEmail account.Email
}

func NewDriver(serviceName string, co shared.ContextOperator, mc mailer.Client, noReplyEmail account.Email) Driver {
	return &driver{serviceName, co, mc, noReplyEmail}
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
		tmplPath = "driver/email/template/otp/otp_template_en.html"
		if err != nil {
			return err
		}
	default:
		subject = "あなたのOTPコード"
		tmplPath = "driver/email/template/otp/otp_template_ja.html"
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
		From:    d.noReplyEmail.ToString(),
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

type inviteMemberTemplateData struct {
	ServiceName   string
	URL           template.URL
	WorkspaceName string
	InviterName   string
}

func (d *driver) SendInvite(ctx context.Context, inviter workspace.Inviter, i *invitation.Invitation) error {
	lang, err := d.co.GetLang(ctx)
	if err != nil {
		return err
	}
	wd := inviter.Workspace().Detail()
	tmplPath := ""
	subject := ""
	switch lang {
	case language.English:
		subject = fmt.Sprintf("%s has invited you to join the %s workspace", inviter.DisplayName().ToString(), wd.Name().ToString())
		tmplPath = "driver/email/template/invite_member/en.html"
		if err != nil {
			return err
		}
	default:
		subject = fmt.Sprintf("%sがあなたを%sワークスペースに招待しました", inviter.DisplayName().ToString(), wd.Name().ToString())
		tmplPath = "driver/email/template/invite_member/ja.html"
		if err != nil {
			return err
		}
	}
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}
	data := inviteMemberTemplateData{
		ServiceName:   d.serviceName,
		URL:           template.URL("https://example.com"),
		WorkspaceName: wd.Name().ToString(),
		InviterName:   inviter.DisplayName().ToString(),
	}
	var body bytes.Buffer
	if err = tmpl.Execute(&body, data); err != nil {
		return err
	}

	conf := mailer.Config{
		From:    d.noReplyEmail.ToString(),
		To:      []string{i.InviteeEmail().ToString()},
		Subject: subject,
		Body:    body.String(),
		Cc:      nil,
		Bcc:     nil,
		ReplyTo: "",
		IsHTML:  true,
	}
	return d.mc.Send(conf)
}

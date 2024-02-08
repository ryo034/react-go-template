package config

import "github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"

const (
	mailPort     Key = "MAIL_PORT"
	mailHost     Key = "MAIL_HOST"
	resendAPIKey Key = "RESEND_API_KEY"
	noReplyEmail Key = "NO_REPLY_EMAIL"
)

func (r *reader) NoReplyEmail() account.Email {
	em := r.fromEnv(noReplyEmail)
	res, err := account.NewEmail(em)
	if err != nil {
		panic(err)
	}
	return res
}

func (r *reader) MailHost() string {
	return r.fromEnv(mailHost)
}

func (r *reader) MailPort() int {
	return r.fromEnvInt(mailPort)
}

func (r *reader) ResendAPIKey() string {
	return r.fromEnv(resendAPIKey)
}

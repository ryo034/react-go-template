package mailer

type Client interface {
	Send(conf Config) error
}

type Config struct {
	From    string
	To      []string
	Subject string
	Body    string
	Cc      []string
	Bcc     []string
	ReplyTo string
	IsHTML  bool
}

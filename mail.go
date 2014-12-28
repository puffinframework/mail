package mail

import (
	"github.com/puffinframework/config"
)

type MailService interface {
	SendMail(to []string, body string) error
}

type mailConfig struct {
	SmtpServer struct {
		User     string
		Password string
	}
}

type mailServiceImpl struct {
}

func NewMailService() MailService {
	cfg := &mailConfig{}
	config.MustReadConfig(cfg)

	return &mailServiceImpl{}
}

func (self *mailServiceImpl) SendMail(to []string, body string) error {
	return nil
}

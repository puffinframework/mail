package mail

import (
	"net/smtp"
	"strconv"
	"strings"

	"github.com/jordan-wright/email"
	"github.com/puffinframework/config"
)

type mailServiceConfig struct {
	SMTP struct {
		Server   string
		Port     int
		Login    string
		Password string
	}
}

type MailService interface {
	Send(e *email.Email) error
}

type implMailService struct {
	smtpAddr string
	smtpAuth smtp.Auth
}

func NewMailService() MailService {
	cfg := &mailServiceConfig{}
	config.MustReadConfig(cfg)

	return &implMailService{
		smtpAddr: strings.Join([]string{cfg.SMTP.Server, strconv.Itoa(cfg.SMTP.Port)}, ":"),
		smtpAuth: smtp.PlainAuth("", cfg.SMTP.Login, cfg.SMTP.Password, cfg.SMTP.Server),
	}
}

func (self *implMailService) Send(e *email.Email) error {
	return e.Send(self.smtpAddr, self.smtpAuth)
}

package mail_test

import (
	"net/smtp"
	"strings"
	"strconv"
	"os"
	"testing"

	"github.com/jordan-wright/email"
	"github.com/puffinframework/config"
	"github.com/stretchr/testify/assert"
)

type mailConfig struct {
	SMTP struct {
		Server   string
		Port     int
		Login    string
		Password string
	}
}

func TestSendMail(t *testing.T) {
	os.Setenv(config.ENV_VAR_NAME, config.MODE_TEST)
	var cfg mailConfig
	config.MustReadConfig(&cfg)

	e := email.NewEmail()
	e.From = cfg.SMTP.Login
	e.To = []string{"dario.freire@gmail.com"}
	e.Subject = "Test Subject"
	e.Text = []byte("Test body text.")
	e.HTML = []byte("Test <i>body</i> <b>html</b>.")

	addr := strings.Join([]string{cfg.SMTP.Server, strconv.Itoa(cfg.SMTP.Port)}, ":")
	auth := smtp.PlainAuth("", cfg.SMTP.Login, cfg.SMTP.Password, cfg.SMTP.Server)
	assert.Nil(t, e.Send(addr, auth))
}

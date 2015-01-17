package mail_test

import (
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jordan-wright/email"
	"github.com/puffinframework/config"
	"github.com/puffinframework/mail"
	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	os.Setenv(config.ENV_VAR_NAME, config.MODE_TEST)
	var cfg mail.MailConfig
	config.MustReadConfig(&cfg)

	now := time.Now().String()

	e := email.NewEmail()
	e.From = cfg.SMTP.Login
	e.To = []string{"dario.freire@gmail.com"}
	e.Subject = "Test Subject"
	e.Text = []byte(now)
	e.HTML = []byte(now)

	addr := strings.Join([]string{cfg.SMTP.Server, strconv.Itoa(cfg.SMTP.Port)}, ":")
	auth := smtp.PlainAuth("", cfg.SMTP.Login, cfg.SMTP.Password, cfg.SMTP.Server)
	assert.Nil(t, e.Send(addr, auth))
}

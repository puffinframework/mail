package mail_test

import (
	"os"
	"testing"
	"time"

	"github.com/jordan-wright/email"
	"github.com/puffinframework/config"
	"github.com/puffinframework/mail"
	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	os.Setenv(config.ENV_VAR_NAME, config.MODE_TEST)
	var cfg mail.MailServiceConfig
	config.MustReadConfig(&cfg)

	ms := mail.NewMailService(cfg)

	e := email.NewEmail()
	e.From = cfg.SMTP.Login
	e.To = []string{"dario.freire@gmail.com"}
	e.Subject = "Test Subject"
	e.HTML = []byte(time.Now().String())

	assert.Nil(t, ms.Send(e))
}

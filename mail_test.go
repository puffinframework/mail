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

	ms := mail.NewMailService()

	e := email.NewEmail()
	e.From = "puffinframework@mailinator.com"
	e.To = []string{"puffinframework@mailinator.com"}
	e.Subject = "Test Subject"
	e.HTML = []byte(time.Now().String())

	assert.Nil(t, ms.Send(e))
}

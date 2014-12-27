package mail

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

func SendMail(to []string, body string) error {
	return nil
}

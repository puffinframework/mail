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

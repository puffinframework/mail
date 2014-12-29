package mail

type MailConfig struct {
	SmtpServer struct {
		Host     string
		Port     int
		User     string
		Password string
	}
}

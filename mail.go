package mail

type MailConfig struct {
	SMTP struct {
		Server   string
		Port     int
		Login    string
		Password string
	}
}

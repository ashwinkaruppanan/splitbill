package utils

import (
	"log"

	"gopkg.in/gomail.v2"
)

type MailService struct {
	host     string
	port     int
	username string
	password string
}

func NewMailService(host string, port int, username string, password string) *MailService {
	return &MailService{
		host,
		port,
		username,
		password,
	}
}

func (m *MailService) MailTo(receiver string, subject string, body string) error {

	// Create a new message
	message := gomail.NewMessage()
	message.SetHeader("From", m.username)
	message.SetHeader("To", receiver)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	// Create a dialer

	dialer := gomail.NewDialer(m.host, m.port, m.username, m.password)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	return nil
}

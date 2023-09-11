package internal

import "fmt"

type EmailSender struct {
	email string
}

func NewEmailSender(email string) *EmailSender {
	return &EmailSender{email: email}
}

func (s *EmailSender) Send(message string) error {
	fmt.Printf("Sending email to: %s message: %s\n", s.email, message)
	return nil
}

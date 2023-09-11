package internal

import "fmt"

type WhatsappSender struct {
	number string
}

func NewWhatsappSender(number string) *WhatsappSender {
	return &WhatsappSender{number: number}
}

func (s *WhatsappSender) Send(message string) error {
	fmt.Printf("Sending whatsapp to: %s message: %s\n", s.number, message)
	return nil
}

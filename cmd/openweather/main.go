package main

import (
	"github.com/fpsbr/openweather/internal"
)

func main() {
	emailSender := internal.NewEmailSender("email@somebody.com")
	whatsappSender := internal.NewWhatsappSender("11 9999-9999")

	lat := -23.588659
	lon := -46.6347239

	wpoller := internal.NewPoller(emailSender, whatsappSender)
	wpoller.Start(lat, lon)
}

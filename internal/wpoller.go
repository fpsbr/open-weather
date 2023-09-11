package internal

import (
	"fmt"
	"log"
	"time"
)

const (
	pollInterval = time.Second * 5
)

type Sender interface {
	Send(message string) error
}

type WPoller struct {
	sender []Sender
}

func NewPoller(s ...Sender) *WPoller {
	return &WPoller{sender: s}
}

func currentTemperature(currentHour int, data *ForecastData) (float64, error) {
	if len(data.Hourly.Temperature2m) < currentHour {
		return 0, ErrInvalidData
	}

	return data.Hourly.Temperature2m[currentHour], nil
}

func (wp *WPoller) sendInformation(data *ForecastData) error {
	for _, s := range wp.sender {
		currentTime := time.Now()
		currentHour := currentTime.Hour()
		t, err := currentTemperature(currentHour, data)

		if err != nil {
			continue
		}

		message := fmt.Sprintf("Hour: <%d> Elevation: <%.2f> Temperature: <%.2f>", currentHour, data.Elevation, t)

		if err := s.Send(message); err != nil {
			fmt.Println("Error sending data...")
		}
	}
	return nil
}

func (wpoller *WPoller) Start(lat float64, lon float64) {
	wf := NewWeatherFetcher()
	ticker := time.Tick(pollInterval)

	for {
		select {
		case <-ticker:
			data, err := wf.GetForecastData(lat, lon)
			if err != nil {
				log.Fatal(err)
			}
			wpoller.sendInformation(data)
		}
	}
}

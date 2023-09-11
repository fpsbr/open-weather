package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	forecast_endpoint = "https://api.open-meteo.com/v1/forecast"
)

type WeatherFetcher struct{}

func NewWeatherFetcher() *WeatherFetcher {
	return &WeatherFetcher{}
}

func (wf *WeatherFetcher) GetForecastData(lat float64, lon float64) (*ForecastData, error) {
	uri := fmt.Sprintf("%s?latitude=%.2f&longitude=%.2f&hourly=temperature_2m", forecast_endpoint, lat, lon)
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data ForecastData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

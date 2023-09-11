package internal

type HourlyData struct {
	Time          []string  `json:"time"`
	Temperature2m []float64 `json:"temperature_2m"`
}

type ForecastData struct {
	Elevation float64    `json: "elevation"`
	Hourly    HourlyData `json:"hourly"`
}

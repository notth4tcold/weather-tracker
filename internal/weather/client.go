package weather

import (
	"encoding/json"
	"net/http"
)

func FetchExternalWeather(apiURL string) (*WeatherAPIResponse, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var w WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		return nil, err
	}
	return &w, nil
}

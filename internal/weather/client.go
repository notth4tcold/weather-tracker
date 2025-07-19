package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client interface {
	FetchExternalWeather(ctx context.Context, city string) (*WeatherAPIResponse, error)
}

type OpenWeatherClient struct {
	apiKey string
}

func NewOpenWeatherClient(apiKey string) *OpenWeatherClient {
	return &OpenWeatherClient{apiKey: apiKey}
}

const baseURL = "https://api.openweathermap.org/data/2.5/weather"

func (c *OpenWeatherClient) FetchExternalWeather(ctx context.Context, city string) (*WeatherAPIResponse, error) {
	cityQuery := url.QueryEscape(city)
	url := fmt.Sprintf("%s?q=%s&units=metric&appid=%s", baseURL, cityQuery, c.apiKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

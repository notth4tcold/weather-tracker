package weather

type WeatherAPIResponse struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temp"`
}

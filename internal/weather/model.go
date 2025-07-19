package weather

type WeatherAPIResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Dt int64 `json:"dt"`
}

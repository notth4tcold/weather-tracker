package service

import (
	"context"
	"log"
	"weather-tracker/internal/db"
	"weather-tracker/internal/models"
	"weather-tracker/internal/weather"
)

type WeatherService struct {
	repo   *db.WeatherRepository
	client weather.Client
}

func NewWeatherService(repo *db.WeatherRepository, client weather.Client) *WeatherService {
	return &WeatherService{
		repo:   repo,
		client: client,
	}
}

func (s *WeatherService) GetAll(ctx context.Context) ([]models.Weather, error) {
	return s.repo.GetAll(ctx)
}

func (s *WeatherService) InsertCity(ctx context.Context, city string, temperature float64) error {
	return s.repo.InsertCity(ctx, city, temperature)
}

func (s *WeatherService) UpdateTemperatureByID(ctx context.Context, id int, temperature float64) error {
	return s.repo.UpdateTemperatureByID(ctx, id, temperature)
}

func (s *WeatherService) DeleteByID(ctx context.Context, id int) error {
	return s.repo.DeleteByID(ctx, id)
}

func (s *WeatherService) GetWeatherByID(ctx context.Context, id int) ([]models.Weather, error) {
	return s.repo.GetWeatherByID(ctx, id)
}

func (s *WeatherService) UpdateAllCitiesWeather(ctx context.Context) error {
	cities, err := s.repo.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, city := range cities {
		weatherData, err := s.client.FetchExternalWeather(ctx, city.City)
		if err != nil {
			log.Printf("Failed to fetch weather for %s: %v", city.City, err)
			continue
		}

		err = s.repo.UpdateTemperatureByID(ctx, city.ID, weatherData.Main.Temp)
		if err != nil {
			log.Printf("Failed to update DB for %s: %v", city.City, err)
		}
	}
	return nil
}

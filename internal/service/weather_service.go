package service

import (
	"weather-tracker/internal/db"
	"weather-tracker/internal/models"
)

func ListWeather() ([]models.Weather, error) {
	return db.GetWeathers()
}

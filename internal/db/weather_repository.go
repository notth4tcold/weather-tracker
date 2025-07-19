package db

import (
	"context"
	"database/sql"
	"weather-tracker/internal/models"
)

type WeatherRepository struct {
	DB *sql.DB
}

func NewWeatherRepository(db *sql.DB) *WeatherRepository {
	return &WeatherRepository{DB: db}
}

func (r *WeatherRepository) GetAll(ctx context.Context) ([]models.Weather, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, city, temperature, timestamp FROM weather ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weathers []models.Weather
	for rows.Next() {
		var w models.Weather
		if err := rows.Scan(&w.ID, &w.City, &w.Temperature, &w.Timestamp); err != nil {
			return nil, err
		}
		weathers = append(weathers, w)
	}

	return weathers, nil
}

func (r *WeatherRepository) InsertCity(ctx context.Context, city string, temperature float64) error {
	_, err := r.DB.ExecContext(ctx, "INSERT INTO weather (city, temperature) VALUES ($1, $2)", city, temperature)
	return err
}

func (r *WeatherRepository) UpdateTemperatureByID(ctx context.Context, id int, temperature float64) error {
	_, err := r.DB.ExecContext(ctx, "UPDATE weather SET temperature = $1, timestamp = NOW() WHERE id = $2", temperature, id)
	return err
}

func (r *WeatherRepository) DeleteByID(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM weather WHERE id = $1", id)
	return err
}

func (r *WeatherRepository) GetWeatherByID(ctx context.Context, id int) ([]models.Weather, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, city, temperature, timestamp FROM weather WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weathers []models.Weather
	for rows.Next() {
		var w models.Weather
		err := rows.Scan(&w.ID, &w.City, &w.Temperature, &w.Timestamp)
		if err != nil {
			return nil, err
		}
		weathers = append(weathers, w)
	}
	return weathers, nil
}

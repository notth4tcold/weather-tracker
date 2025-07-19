package db

import (
	"weather-tracker/internal/models"
)

func GetWeathers() ([]models.Weather, error) {
	rows, err := DB.Query("SELECT id, city, temperature FROM weather")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weathers []models.Weather
	for rows.Next() {
		var w models.Weather
		if err := rows.Scan(&w.ID, &w.City, &w.Temperature); err != nil {
			return nil, err
		}
		weathers = append(weathers, w)
	}

	return weathers, nil
}

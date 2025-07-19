package main

import (
	"log"
	"weather-tracker/config"
	"weather-tracker/internal/api"
	"weather-tracker/internal/background"
	"weather-tracker/internal/db"
	"weather-tracker/internal/service"
	"weather-tracker/internal/weather"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	sqlDB, err := db.Init(cfg)
	if err != nil {
		log.Fatalf("DB init failed: %v", err)
	}
	defer sqlDB.Close()
	log.Println("Connected to DB successfully")

	repo := db.NewWeatherRepository(sqlDB)
	client := weather.NewOpenWeatherClient(cfg.OpenWeatherMapAPIKey)
	service := service.NewWeatherService(repo, client)

	r := gin.Default()
	api.RegisterRoutes(r, service)

	stopCh := make(chan struct{})
	go background.StartUpdater(service, stopCh)
	defer func() {
		close(stopCh)
	}()

	log.Printf("Starting server on :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

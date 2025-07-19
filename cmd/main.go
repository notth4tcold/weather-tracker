package main

import (
	"log"
	"weather-tracker/config"
	"weather-tracker/internal/api"
	"weather-tracker/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db.Init(cfg)

	r := gin.Default()
	api.RegisterRoutes(r)

	log.Printf("Starting server on :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

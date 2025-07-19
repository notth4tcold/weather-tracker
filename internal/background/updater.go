package background

import (
	"context"
	"log"
	"time"
	"weather-tracker/internal/service"
)

func StartUpdater(service *service.WeatherService, stopCh <-chan struct{}) {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Running periodic update...")

			if err := service.UpdateAllCitiesWeather(context.Background()); err != nil {
				log.Printf("Weather update failed: %v", err)
			}
		case <-stopCh:
			log.Println("Weather updater stopped")
			return
		}
	}
}

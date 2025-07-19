package api

import (
	"net/http"
	"strconv"
	"weather-tracker/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, service *service.WeatherService) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/cities", func(c *gin.Context) {
		cities, err := service.GetAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cities)
	})

	r.POST("/cities", func(c *gin.Context) {
		var req struct {
			City        string  `json:"city"`
			Temperature float64 `json:"temperature"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		if err := service.InsertCity(c.Request.Context(), req.City, req.Temperature); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not insert"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "city added"})
	})

	r.PUT("/cities/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		var body struct {
			Temperature float64 `json:"temperature"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
			return
		}
		if err := service.UpdateTemperatureByID(c.Request.Context(), id, body.Temperature); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update temperature"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "temperature updated"})
	})

	r.DELETE("/cities/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		if err := service.DeleteByID(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	})

	r.GET("/cities/:id/weather", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		w, err := service.GetWeatherByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "city not found"})
			return
		}
		c.JSON(http.StatusOK, w)
	})
}

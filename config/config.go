package config

import (
	"log"
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	ServerPort string
}

func LoadConfig() Config {
	getEnvOrFail := func(key string) string {
		val := os.Getenv(key)
		if val == "" {
			log.Fatalf("Environment variable %s is required", key)
		}
		return val
	}

	return Config{
		DBUser:     getEnvOrFail("DB_USER"),
		DBPassword: getEnvOrFail("DB_PASSWORD"),
		DBHost:     getEnvOrFail("DB_HOST"),
		DBPort:     getEnvOrFail("DB_PORT"),
		DBName:     getEnvOrFail("DB_NAME"),
		ServerPort: getEnvOrFail("SERVER_PORT"),
	}
}
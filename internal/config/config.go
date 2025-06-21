package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_DRIVER    string
	DB_DSN       string
	PORT         string
	VAT_RATE     float64
	SECRET_KEY   string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	vatRate := 0.2
	if envVat := os.Getenv("VAT_RATE"); envVat != "" {
		if rate, err := strconv.ParseFloat(envVat, 64); err == nil && rate >= 0 && rate <= 1 {
			vatRate = rate
		}
	}

	return &Config{
		DB_DSN: fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		),
		DB_DRIVER:    os.Getenv("DB_DRIVER"),
		PORT:         os.Getenv("PORT"),
		VAT_RATE:     vatRate,
		SECRET_KEY:   os.Getenv("SECRET_KEY"),
	}
}

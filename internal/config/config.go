package config

import (
	"fmt"
	"os"
)

type Config struct {
	DB_DRIVER string
	DB_DSN    string
	PORT      string
}

func Load() *Config {
	return &Config{
		DB_DSN: fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		),
		DB_DRIVER: os.Getenv("DB_DRIVER"),
		PORT:      os.Getenv("PORT"),
	}
}

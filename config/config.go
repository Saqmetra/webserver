package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config структура с настройками
type Config struct {
	Port string
}

// LoadConfig загружает конфигурацию из .env
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Файл .env не найден, используем значения по умолчанию")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{Port: port}
}

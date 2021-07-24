package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Auth   AuthConfig
	DB     DBConfig
	HTTP   HTTPConfig
	Sentry SentryConfig
}

// NewConfig func
func NewConfig() *Config {
	switch godotenv.Load() {
	case godotenv.Load("../.env"):
		log.Println("Error loading .env file")
	}

	return &Config{
		Auth:   LoadAuthConfig(),
		DB:     LoadDBConfig(),
		HTTP:   LoadHTTPConfig(),
		Sentry: LoadSentryConfig(),
	}
}

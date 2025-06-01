package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB DbConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No envs found. Starting with default values")
	}

	return &Config{
		DB: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
	}
}

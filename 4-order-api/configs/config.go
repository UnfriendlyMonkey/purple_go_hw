package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB  DbConfig
	JWT JWTConfig
}

type DbConfig struct {
	Dsn string
}

type JWTConfig struct {
	Secret string
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
		JWT: JWTConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}

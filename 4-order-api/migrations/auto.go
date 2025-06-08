package main

import (
	"go/hw/4-order-api/internal/product"
	"go/hw/4-order-api/internal/user"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db.AutoMigrate(&product.Product{}, &user.User{})

	log.Println("Database migrated")
}

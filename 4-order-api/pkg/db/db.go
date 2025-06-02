package db

import (
	"go/hw/4-order-api/configs"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDB(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.DB.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &Db{db}
}

package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:Querentino1!@tcp(localhost:3307)/banco?charset=utf8mb4&parseTime=True&loc=Local"
)

func NewDatabaseConnection() *gorm.DB {

	log.Println("Initializing database...")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Error during database initialization!")
	}

	log.Println("Successfully initialized database!")

	return db

}

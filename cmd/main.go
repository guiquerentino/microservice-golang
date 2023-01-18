package main

import (
	"log"
	"microservice-golang/api/entities"
	"microservice-golang/api/routes"
	"microservice-golang/config"
)

var (
	db = config.NewDatabaseConnection()
)

func init() {
	log.Println("Migrating database models...")

	db.AutoMigrate(&entities.Album{}, &entities.Music{})
	log.Println("Successfully migrated!")
}

func main() {

	route := routes.SetupRoutes(db)
	route.Run(":8080")

}

package routes

import (
	"encoding/json"
	"fmt"
	"microservice-golang/api/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

	routes := gin.Default()

	routes.POST("api/album", func(context *gin.Context) {
		var request entities.AlbumDTO

		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(400, gin.H{
				"message": "Error processing your request",
			})
		} else {

			var objetoDB entities.Album
			objetoDB.Author = request.Author
			objetoDB.Genre = request.Genre
			objetoDB.Name = request.Name

			musicsJson, err := json.Marshal(request.Musics)

			if err != nil {
				context.JSON(500, gin.H{
					"message": "Error during JSON Serialization!",
				})
			} else {
				objetoDB.Musics = string(musicsJson)
			}

			db.Create(&objetoDB)

			context.JSON(200, gin.H{
				"message": "Album successfully added!",
			})
		}
	})

	routes.GET("/api/album/:id", func(context *gin.Context) {
		var object entities.Album

		id := context.Param("id")

		err := db.First(&object, id)

		if err.Error != nil {
			context.JSON(404, gin.H{
				"message": "Album not found!",
			})
		} else {

			var response entities.AlbumDTO
			json.Unmarshal([]byte(object.Musics), &response.Musics)
			response.Name = object.Name
			response.Genre = object.Genre
			response.Author = object.Author

			context.JSON(200, response)
		}

	})

	routes.PUT("api/album/:id", func(context *gin.Context) {
		var request entities.AlbumDTO
		var dbObject entities.Album

		id := context.Param("id")

		fmt.Println(id)

		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(400, gin.H{
				"message": "Error processing your request",
			})
		} else {

			if err := db.First(&dbObject, id); err.Error != nil {
				context.JSON(404, gin.H{
					"message":    "Album not found!",
					"stacktrace": err.Name(),
				})
			} else {

				dbObject.Genre = request.Genre
				dbObject.Author = request.Author
				dbObject.Name = request.Name

				musicsJson, err := json.Marshal(request.Musics)

				if err != nil {
					context.JSON(500, gin.H{
						"message": "Error during JSON Serialization!",
					})
				} else {
					dbObject.Musics = string(musicsJson)
				}

				db.Save(&dbObject)

				context.JSON(200, gin.H{
					"message": "Album successfully updated!",
				})
			}
		}
	})

	routes.DELETE("api/album/:id", func(context *gin.Context) {
		var request entities.Album

		id := context.Param("id")

		err := db.Delete(&request, id)

		if err.Error != nil {
			context.JSON(404, gin.H{
				"message":    "Album not found!",
				"stacktrace": err.Name(),
			})
		} else {

			context.JSON(200, gin.H{
				"mensage": "Album successfully deleted!",
			})
		}

	})
	return routes
}

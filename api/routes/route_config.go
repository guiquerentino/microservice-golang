package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"microservice-golang/api/entities"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

	routes := gin.Default()

	routes.POST("api/album", func(context *gin.Context) {
		var request entities.Album

		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(400, gin.H{
				"message": "Erro processing your request",
			})
		} else {

			db.Create(&request)

			context.JSON(200, gin.H{
				"message": "Album successfully added!",
			})
		}
	})

	routes.GET("/api/album/:id", func(context *gin.Context) {
		var object entities.Album

		id := context.Param("id")

		err := db.First(&object, id)

		if err != nil {
			context.JSON(404, gin.H{
				"message":    "Album not found!",
				"stacktrace": err.Error,
			})
		} else {
			context.JSON(200, object)
		}

	})

	routes.PUT("api/album/:id", func(context *gin.Context) {
		var request entities.Album
		var dbObject entities.Album

		id := context.Param("id")

		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(400, gin.H{
				"message": "Error processing your request",
			})
		} else {

			err := db.First(&dbObject, id)

			if err != nil {
				context.JSON(404, gin.H{
					"message":    "Album not found!",
					"stacktrace": err.Name(),
				})
			} else {

				dbObject.Genre = request.Genre
				dbObject.Author = request.Author
				dbObject.Name = request.Name
				//dbObject.Musics = request.Musics

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

		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(400, gin.H{
				"message": "Erro processing your request",
			})
		} else {

			err := db.Delete(&request, id)

			if err != nil {
				context.JSON(404, gin.H{
					"message":    "Album not found!",
					"stacktrace": err.Name(),
				})
			} else {

				context.JSON(200, gin.H{
					"mensagem": "Album successfully deleted!",
				})
			}
		}
	})

	return routes

}

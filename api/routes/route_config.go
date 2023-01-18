package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

	routes := gin.Default()

	routes.POST("api/album", func(context *gin.Context) {

	})

	routes.GET("/api/album/:id", func(context *gin.Context) {

	})

	routes.PUT("api/album/:id", func(context *gin.Context) {

	})

	routes.DELETE("api/album/:id", func(context *gin.Context) {

	})

	return routes

}

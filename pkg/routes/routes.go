package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/narekk1202/pictures-app/pkg/handlers"
)

func SetRoutes(router *gin.Engine) {
	router.GET("/pictures", handlers.GetPictures)
	router.POST("/pictures", handlers.AddPicture)
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narekk1202/pictures-app/pkg/database"
	"github.com/narekk1202/pictures-app/pkg/models"
)

func GetPictures(ctx *gin.Context) {
	var picturesFromDB []models.Pictures
	DB := database.GetDB()

	if err := DB.Table("pictures.pictures").Find(&picturesFromDB).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pictures"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"pictures": picturesFromDB,
	})
}

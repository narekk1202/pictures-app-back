package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narekk1202/pictures-app/pkg/database"
	"github.com/narekk1202/pictures-app/pkg/models"
)

func AddPicture(ctx *gin.Context) {
	var newPicture models.AddPicture
	if err := ctx.ShouldBindJSON(&newPicture); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB := database.GetDB()

	picture := models.Pictures{
		PictureUrl:   newPicture.PictureUrl,
		UploaderName: newPicture.UploaderName,
	}

	if err := DB.Table("pictures.pictures").Create(&picture).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	DB.Table("pictures.pictures").Create(&picture)
	ctx.JSON(http.StatusOK, gin.H{"message": "Picture added successfully"})
}

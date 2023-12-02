package models

import "gorm.io/gorm"

type Pictures struct {
	// ID           uint   `gorm:"primaryKey" json:"id"`
	gorm.Model
	PictureUrl   string `json:"picture_url"`
	UploaderName string `json:"uploader_name"`
}

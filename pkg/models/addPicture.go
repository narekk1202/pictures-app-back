package models

type AddPicture struct {
	PictureUrl   string `json:"picture_url"`
	UploaderName string `json:"uploader_name"`
}
package model

type CreateFileRequest struct {
	Name   string `json:"file" validate:"required"`
	Gambar Image  `json:"gambar" validate:"required,dive,required"`
}

type Image struct {
	Filename    string `validate:"required"`
	Ext         string `validate:"required,image_validation"`
	ContentType string `validate:"required"`
	Bytes       int32  `validate:"required,max=2048"`
}

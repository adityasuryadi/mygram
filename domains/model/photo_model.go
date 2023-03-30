package model

import (
	"time"

	"github.com/google/uuid"
)

type CreatePhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	PhotoUrl string `json:"photo_url" validate:"required"`
	Caption  string `json:"caption"`
	Email string
}

type CreatePhotoResponse struct {
	Id uuid.UUID `json:"id"`
	PhotoUrl string `json:"photo_url"`
	Caption string `json:"caption"`
	Title string `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
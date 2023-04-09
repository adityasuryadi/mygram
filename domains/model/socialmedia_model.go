package model

import "time"

type CreateSocialmediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialmediaUrl string `json:"social_media_url" validate:"required"`
	Email          string
}

type SocialmediaResponse struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	SocialmediaUrl string `json:"social_media_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt 	   time.Time `json:"updated_at"`
}

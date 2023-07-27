package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductCreateRequest struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type ProductResponse struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Stock     int       `json:"stock"`
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
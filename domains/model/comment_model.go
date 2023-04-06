package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateCommentRequest struct {
	PhotoId uuid.UUID `json:"photo_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type UpdateCommentRequest struct {
	PhotoId uuid.UUID `json:"photo_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type CommentResponse struct {
	Id uuid.UUID `json:"id"`
	PhotoId uuid.UUID `json:"photo_id"`
	Message string `json:"messgae"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
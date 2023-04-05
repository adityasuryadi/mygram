package model

import "github.com/google/uuid"

type CreateCommentRequest struct {
	PhotoId uuid.UUID `json:"photo_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}
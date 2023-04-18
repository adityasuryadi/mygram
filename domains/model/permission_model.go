package model

import "time"

type CreatePermissionRequest struct {
	Name string `json:"name"`
}

type PermissionResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}
package models

import (
	"time"

	"github.com/google/uuid"
)

const USER = "user"

type User struct {
	UserID       uuid.UUID `json:"user_id"`
	UserName     string    `json:"user_name" validate:"required, min=2, max=100"`
	Email        string    `json:"email" validate:"eamil, required"`
	Password     string    `json:"password" validate:"required, min=8"`
	Token        *string   `json:"token"`
	RefreshToken *string   `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

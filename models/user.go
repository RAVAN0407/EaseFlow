package models

import (
	"time"

	"github.com/google/uuid"
)

const USER = "user"

type User struct {
	UserID       uuid.UUID `json:"userId"`
	UserName     string    `json:"username" validate: "required,min=2,max=100"`
	Email        string    `json:"email" validate:"email,required"`
	Password     string    `json:"password" validate:"required,min=8"`
	Token        *string   `json:"token"`
	RefreshToken *string   `json:"refreshToken"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

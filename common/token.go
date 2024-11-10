package common

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type UserClaims struct {
	UserID   uuid.UUID
	UserName string
	Email    string
	Password string
	jwt.StandardClaims
}

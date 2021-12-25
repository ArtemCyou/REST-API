package service

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(name, password string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name
}
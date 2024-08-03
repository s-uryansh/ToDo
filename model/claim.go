package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

package jwt

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

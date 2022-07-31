package structs

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Email   string `json:"email"`
	IsAdmin int    `json:"is_admin"`
	jwt.StandardClaims
}

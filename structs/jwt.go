package structs

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Id      int    `json:"id"`
	Email   string `json:"email"`
	IsAdmin int    `json:"is_admin"`
	jwt.StandardClaims
}

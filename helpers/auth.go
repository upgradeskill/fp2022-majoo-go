package helpers

import (
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Auth(c echo.Context) (jwt.MapClaims, bool) {
	authorization := c.Request().Header.Get("Authorization")
	tokenStr := strings.Split(authorization, "Bearer ")[1]

	hmacSecretString := os.Getenv("JWT_SECRET")
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

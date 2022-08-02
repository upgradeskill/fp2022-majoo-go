package controllers

import (
	"net/http"
	"os"
	"time"

	"mini-pos/configs"
	model "mini-pos/models"
	"mini-pos/structs"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := model.GetOneUserByEmail(email) // method get by email
	response := new(structs.Response)

	if err != nil {
		response.Message = "User tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {

		if user.Password != password {
			response.Message = "Password salah"
			return c.JSON(http.StatusBadRequest, response)
		}

		// Set custom claims
		claims := &structs.JwtCustomClaims{
			Id:      user.Id,
			Email:   user.Email,
			IsAdmin: user.IsAdmin,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"messsage": "Berhasil login",
			"data":     user,
			"token":    t,
		})
	}
}

func RedisPing(c echo.Context) error {
	client := configs.RedisClient()

	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, pong)
}

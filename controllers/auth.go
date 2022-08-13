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

func Register(c echo.Context) error {
	user := new(structs.Users)
	c.Bind(user)

	response := new(structs.Response)
	checkUser, _ := model.GetOneUserByEmail(user.Email) // method get by email

	if checkUser.Email == user.Email {
		response.Message = "Email sudah pernah terdaftar"
		return c.JSON(http.StatusBadRequest, response)
	}

	if model.CreateUser(user) != nil { // method create user
		response.Message = "Gagal create data"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Message = "Sukses create data"
		response.Data = user
		return c.JSON(http.StatusOK, response)
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

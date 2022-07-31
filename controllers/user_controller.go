package controllers

import (
	"fmt"
	"mini-pos/helpers"
	model "mini-pos/models"
	"mini-pos/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Profile(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return err
	}

	token, _ := helpers.Auth(cookie.Value)

	response := new(structs.Response)
	response.Status = 200
	response.Message = "Sukses melihat data"
	response.Data = token

	return c.JSON(http.StatusOK, response)
}

func UserList(c echo.Context) error {
	response := new(structs.Response)
	users, err := model.GetAllUser(c.QueryParam("q")) // method get all

	if err != nil {
		response.Status = 400
		response.Message = "Gagal melihat data"
		return c.JSON(http.StatusBadRequest, response)
	} else {
		response.Status = 200
		response.Message = "Sukses melihat data"
		response.Data = users
		return c.JSON(http.StatusOK, response)
	}
}

func UserStore(c echo.Context) error {
	user := new(structs.Users)
	c.Bind(user)
	contentType := c.Request().Header.Get("Content-type")
	if contentType == "application/json" {
		fmt.Println("Request dari json")
	}
	response := new(structs.Response)

	if model.CreateUser(user) != nil { // method create user
		response.Status = 500
		response.Message = "Gagal create data"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Status = 200
		response.Message = "Sukses create data"
		response.Data = *user
		return c.JSON(http.StatusOK, response)
	}
}

func UserShow(c echo.Context) error {
	user, err := model.GetOneUserById(c.Param("id")) // method get by email
	response := new(structs.Response)

	if err != nil {
		response.Status = 404
		response.Message = "User tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {
		response.Status = 200
		response.Message = "Sukses melihat data"
		response.Data = user
		return c.JSON(http.StatusOK, response)
	}
}

func UserUpdate(c echo.Context) error {
	user := new(structs.Users)
	c.Bind(user)
	response := new(structs.Response)
	if model.UpdateUser(c.Param("id"), user) != nil { // method update user
		response.Status = 500
		response.Message = "Gagal update data"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Status = 200
		response.Message = "Sukses update data"
		response.Data = *user
		return c.JSON(http.StatusOK, response)
	}
}

func UserDelete(c echo.Context) error {
	user, _ := model.GetOneUserById(c.Param("id"))
	response := new(structs.Response)

	if model.DeleteUser(&user) != nil {
		response.Status = 404
		response.Message = "User tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {
		response.Status = 200
		response.Message = "Sukses menghapus data user"
		return c.JSON(http.StatusOK, response)
	}
}

package controllers

import (
	"fmt"
	model "mini-pos/models"
	"mini-pos/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserList(c echo.Context) error {
	response := new(structs.Response)
	users, err := model.GetAll(c.QueryParam("keywords")) // method get all

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
	user := new(model.Users)
	c.Bind(user)
	contentType := c.Request().Header.Get("Content-type")
	if contentType == "application/json" {
		fmt.Println("Request dari json")
	}
	response := new(structs.Response)
	if user.CreateUser() != nil { // method create user
		response.Status = 500
		response.Message = "Gagal create data"
		return c.JSON(http.StatusBadRequest, response)
	} else {
		response.Status = 200
		response.Message = "Sukses create data"
		response.Data = *user
		return c.JSON(http.StatusOK, response)
	}
}

func UserShow(c echo.Context) error {
	user, err := model.GetOneById(c.Param("id")) // method get by email
	response := new(structs.Response)

	if err != nil {
		response.Status = 404
		response.Message = "User tidak ditemukan"
		return c.JSON(http.StatusBadRequest, response)
	} else {
		response.Status = 200
		response.Message = "Sukses melihat data"
		response.Data = user
		return c.JSON(http.StatusOK, response)
	}
}

func UserUpdate(c echo.Context) error {
	user := new(model.Users)
	c.Bind(user)
	response := new(structs.Response)
	if user.UpdateUser(c.Param("id")) != nil { // method update user
		response.Status = 500
		response.Message = "Gagal update data"
		return c.JSON(http.StatusBadRequest, response)
	} else {
		response.Status = 200
		response.Message = "Sukses update data"
		response.Data = *user
		return c.JSON(http.StatusOK, response)
	}
}

func UserDelete(c echo.Context) error {
	user, _ := model.GetOneById(c.Param("id"))
	response := new(structs.Response)

	if user.DeleteUser() != nil {
		response.Status = 404
		response.Message = "User tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {
		response.Status = 200
		response.Message = "Sukses menghapus data user"
		return c.JSON(http.StatusOK, response)
	}
}

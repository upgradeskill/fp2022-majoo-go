package controllers

import (
	"fmt"
	"mini-pos/helpers"
	model "mini-pos/models"
	"mini-pos/structs"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UserList(c echo.Context) error {
	response := new(structs.Response)

	auth, _ := helpers.Auth(c)
	userId := fmt.Sprint(auth["id"])
	getOutletUserByUserId, err := model.GetOutletUserByUserId(userId)

	if err != nil {
		response.Message = "Kamu belum memiliki outlet"
		return c.JSON(http.StatusBadRequest, response)
	}

	outletsId := make([]interface{}, len(getOutletUserByUserId))
	for i := 0; i < len(getOutletUserByUserId); i++ {
		outletsId[i] = getOutletUserByUserId[i].OutletId
	}

	getOutletUserByOutletId, err := model.GetOutletUserByOutletId(outletsId)
	if err != nil {
		response.Message = "Gagal get outlet user"
		return c.JSON(http.StatusBadRequest, response)
	}

	usersId := make([]interface{}, len(getOutletUserByOutletId))
	for i := 0; i < len(getOutletUserByOutletId); i++ {
		usersId[i] = getOutletUserByOutletId[i].UserId
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	users, err := model.GetAllUser(c.QueryParam("q"), usersId, limit, offset) // method get all

	if err != nil {
		response.Message = "Gagal melihat data"
		return c.JSON(http.StatusBadRequest, response)
	} else {
		response.Message = "Sukses melihat data"
		response.Data = users
		return c.JSON(http.StatusOK, response)
	}
}

func UserStore(c echo.Context) error {
	user := new(structs.Users)
	c.Bind(user)

	response := new(structs.ResponsePagination)
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

func UserShow(c echo.Context) error {
	user, err := model.GetOneUserById(c.Param("id")) // method get by email
	response := new(structs.Response)

	if err != nil {
		response.Message = "User tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {
		response.Message = "Sukses melihat data"
		response.Data = user
		return c.JSON(http.StatusOK, response)
	}
}

func UserUpdate(c echo.Context) error {
	user := new(structs.Users)
	c.Bind(user)
	response := new(structs.Response)

	if user.Email != "" {
		response.Message = "Email tidak boleh diupdate"
		return c.JSON(http.StatusInternalServerError, response)
	}

	if model.UpdateUser(c.Param("id"), user) != nil { // method update user
		response.Message = "Gagal update data"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Message = "Sukses update data"
		response.Data = *user
		return c.JSON(http.StatusOK, response)
	}
}

func UserDelete(c echo.Context) error {
	user, _ := model.GetOneUserById(c.Param("id"))
	response := new(structs.Response)

	if model.DeleteUser(&user) != nil {
		response.Message = "User tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {
		response.Message = "Sukses menghapus data user"
		return c.JSON(http.StatusOK, response)
	}
}

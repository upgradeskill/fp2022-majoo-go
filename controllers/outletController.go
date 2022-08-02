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

func OutletList(c echo.Context) error {
	response := new(structs.ResponsePagination)

	auth, _ := helpers.Auth(c)
	userId := fmt.Sprint(auth["id"])
	outletUsers, err := model.GetOutletUserByUserId(userId)
	outletsId := make([]interface{}, len(outletUsers))

	for i := 0; i < len(outletUsers); i++ {
		outletsId[i] = outletUsers[i].OutletId
	}

	fmt.Println("user", outletsId)

	if err != nil {
		response.Message = "Kamu belum memiliki outlet"
		return c.JSON(http.StatusBadRequest, response)
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	outlets, err := model.GetAllOutlet(c.QueryParam("q"), outletsId, limit, offset) // method get all

	if err != nil {
		response.Message = "Gagal melihat data"
		return c.JSON(http.StatusBadRequest, response)
	} else {
		response.Message = "Sukses melihat data"
		response.Data = outlets
		response.Limit = limit
		response.Offset = offset
		return c.JSON(http.StatusOK, response)
	}
}

func OutletStore(c echo.Context) error {
	outlet := new(structs.Outlets)
	c.Bind(outlet)

	response := new(structs.Response)

	if model.CreateOutlet(outlet) != nil { // method create outlet
		response.Message = "Gagal create data"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		auth, isAuth := helpers.Auth(c)
		if !isAuth {
			response.Message = "Token tidak valid"
			return c.JSON(http.StatusUnauthorized, response)
		}

		userId, err := strconv.Atoi(fmt.Sprint(auth["id"]))
		if err != nil {
			response.Message = "Convert id gagal"
			return c.JSON(http.StatusInternalServerError, response)
		}

		outletUser := new(structs.OutletUsers)
		outletUser.UserId = userId
		outletUser.OutletId = outlet.Id

		if model.CreateOutletUser(outletUser) != nil { // method create outlet
			response.Message = "Gagal create outlet user"
			return c.JSON(http.StatusInternalServerError, response)
		} else {
			response.Message = "Sukses create data"
			response.Data = outlet
			return c.JSON(http.StatusOK, response)
		}
	}
}

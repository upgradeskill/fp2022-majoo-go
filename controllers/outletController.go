package controllers

import (
	"mini-pos/helpers"
	model "mini-pos/models"
	"mini-pos/structs"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func OutletList(c echo.Context) error {
	response := new(structs.ResponsePagination)
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	outlets, err := model.GetAllOutlet(c.QueryParam("q"), limit, offset) // method get all

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
		cookie, err := c.Cookie("token")
		if err != nil {
			response.Message = "Cookie key tidak tersedia"
			return c.JSON(http.StatusInternalServerError, response)
		}

		user, isAuth := helpers.Auth(cookie.Value)

		if !isAuth {
			response.Message = "Token tidak valid"
			return c.JSON(http.StatusUnauthorized, response)
		}

		userId, ok := user["id"]
		if !ok {
			response.Message = "Convert id gagal"
			return c.JSON(http.StatusInternalServerError, response)
		}

		outletUser := new(structs.OutletUsers)
		outletUser.UserId = int(userId.(float64))
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

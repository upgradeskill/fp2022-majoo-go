package controllers

import (
	model "mini-pos/models"
	"mini-pos/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CategoryList(c echo.Context) error {
	response := new(structs.Response)
	users, err := model.GetAllCategory(c.QueryParam("keywords")) // method get all

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

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

func CategoryList(c echo.Context) error {
	response := new(structs.ResponsePagination)

	auth, _ := helpers.Auth(c)
	userId := fmt.Sprint(auth["id"])
	outletUsers, err := model.GetOutletUserByUserId(userId)

	if err != nil {
		response.Message = "Kamu belum memiliki outlet"
		return c.JSON(http.StatusBadRequest, response)
	}

	outletsId := make([]interface{}, len(outletUsers))
	for i := 0; i < len(outletUsers); i++ {
		outletsId[i] = outletUsers[i].OutletId
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	categories, err := model.GetAllCategory(c.QueryParam("q"), outletsId, limit, offset) // method get all

	if err != nil {
		response.Message = "Gagal melihat data"
		return c.JSON(http.StatusBadRequest, response)
	} else {
		response.Message = "Sukses melihat data"
		response.Data = categories
		response.Limit = limit
		response.Offset = offset
		return c.JSON(http.StatusOK, response)
	}
}

func CategoryStore(c echo.Context) error {
	category := new(structs.Categories)
	c.Bind(category)
	contentType := c.Request().Header.Get("Content-type")
	if contentType == "application/json" {
		fmt.Println("Request dari json")
	}
	response := new(structs.Response)

	if model.CreateCategory(category) != nil { // method create category
		response.Message = "Gagal create data"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Message = "Sukses create data"
		// response.Data = *category
		return c.JSON(http.StatusOK, response)
	}
}

func CategoryShow(c echo.Context) error {
	category, err := model.GetOneCategoryById(c.Param("id")) // method get by email
	response := new(structs.Response)

	if err != nil {
		response.Message = "Data tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {
		response.Message = "Sukses melihat data"
		response.Data = category
		return c.JSON(http.StatusOK, response)
	}
}

func CategoryUpdate(c echo.Context) error {
	category := new(structs.Categories)
	c.Bind(category)
	response := new(structs.Response)
	if model.UpdateCategory(c.Param("id"), category) != nil { // method update category
		response.Message = "Gagal update data"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Message = "Sukses update data"
		// response.Data = *categorys
		return c.JSON(http.StatusOK, response)
	}
}

func CategoryDelete(c echo.Context) error {
	category, _ := model.GetOneCategoryById(c.Param("id"))
	response := new(structs.Response)

	if model.DeleteCategory(&category) != nil {
		response.Message = "data tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {
		response.Message = "Sukses menghapus data"
		return c.JSON(http.StatusOK, response)
	}
}

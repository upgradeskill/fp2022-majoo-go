package controllers

import (
	"fmt"
	model "mini-pos/models"
	"mini-pos/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CategoryList(c echo.Context) error {
	response := new(structs.Response)
	categories, err := model.GetAllCategory(c.QueryParam("q")) // method get all

	if err != nil {
		response.Status = 400
		response.Message = "Gagal melihat data"
		return c.JSON(http.StatusBadRequest, response)
	} else {
		response.Status = 200
		response.Message = "Sukses melihat data"
		response.Data = categories
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
		response.Status = 500
		response.Message = "Gagal create data"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Status = 200
		response.Message = "Sukses create data"
		response.Data = *category
		return c.JSON(http.StatusOK, response)
	}
}

func CategoryShow(c echo.Context) error {
	category, err := model.GetOneCategoryById(c.Param("id")) // method get by email
	response := new(structs.Response)

	if err != nil {
		response.Status = 404
		response.Message = "Data tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {
		response.Status = 200
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
		response.Status = 500
		response.Message = "Gagal update data"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Status = 200
		response.Message = "Sukses update data"
		response.Data = *category
		return c.JSON(http.StatusOK, response)
	}
}

func CategoryDelete(c echo.Context) error {
	category, _ := model.GetOneCategoryById(c.Param("id"))
	response := new(structs.Response)

	if model.DeleteCategory(&category) != nil {
		response.Status = 404
		response.Message = "data tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {
		response.Status = 200
		response.Message = "Sukses menghapus data"
		return c.JSON(http.StatusOK, response)
	}
}

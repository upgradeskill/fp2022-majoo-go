package controllers

import (
	"fmt"
	model "mini-pos/models"
	"mini-pos/structs"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ProductList(c echo.Context) error {
	response := new(structs.ResponsePagination)

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	categories, err := model.GetAllProduct(c.QueryParam("q"), limit, offset) // method get all

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

func ProductStore(c echo.Context) error {
	product := new(structs.Products)
	c.Bind(product)
	contentType := c.Request().Header.Get("Content-type")
	if contentType == "application/json" {
		fmt.Println("Request dari json")
	}
	response := new(structs.Response)

	if model.CreateProduct(product) != nil { // method create product
		response.Message = "Gagal create data"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Message = "Sukses create data"
		// response.Data = *product
		return c.JSON(http.StatusOK, response)
	}
}

func ProductShow(c echo.Context) error {
	product, err := model.GetOneProductById(c.Param("id")) // method get by email
	response := new(structs.Response)

	if err != nil {
		response.Message = "Data tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {
		response.Message = "Sukses melihat data"
		response.Data = product
		return c.JSON(http.StatusOK, response)
	}
}

func ProductUpdate(c echo.Context) error {
	product := new(structs.Products)
	c.Bind(product)
	response := new(structs.Response)
	if model.UpdateProduct(c.Param("id"), product) != nil { // method update product
		response.Message = "Gagal update data"
		return c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Message = "Sukses update data"
		// response.Data = *product
		return c.JSON(http.StatusOK, response)
	}
}

func ProductDelete(c echo.Context) error {
	product, _ := model.GetOneProductById(c.Param("id"))
	response := new(structs.Response)

	if model.DeleteProduct(&product) != nil {
		response.Message = "data tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {
		response.Message = "Sukses menghapus data"
		return c.JSON(http.StatusOK, response)
	}
}

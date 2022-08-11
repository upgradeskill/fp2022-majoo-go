package controllers

import (
	"encoding/json"
	"fmt"
	"mini-pos/helpers"
	model "mini-pos/models"
	"mini-pos/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func TransactionList(c echo.Context) error {
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

	fmt.Println("outlets", outletsId)

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	transactions, err := model.GetAllTransaction(c.QueryParam("q"), outletsId, limit, offset) // method get all

	if err != nil {
		response.Message = "Gagal melihat data"
		return c.JSON(http.StatusBadRequest, response)
	} else {
		response.Message = "Sukses melihat data"
		response.Data = transactions
		response.Limit = limit
		response.Offset = offset
		return c.JSON(http.StatusOK, response)
	}
}

func TransactionStore(c echo.Context) error {

	requestBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&requestBody)

	response := new(structs.Response)

	if err != nil {
		response.Message = "Convert request body"
		return c.JSON(http.StatusInternalServerError, response)
	} else {

		outletId, err := strconv.Atoi(fmt.Sprint(requestBody["outlet_id"]))
		if err != nil {
			response.Message = "Convert outlet id gagal"
			return c.JSON(http.StatusInternalServerError, response)
		}

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

		transaction := new(structs.Transactions)
		transaction.OutletId = outletId
		transaction.Code = strconv.Itoa(int(time.Now().Unix()))
		transaction.CustomerName = fmt.Sprint(requestBody["customer_name"])
		transaction.CreatedBy = userId

		if model.CreateTransaction(transaction) != nil { // method create outlet
			response.Message = "Gagal create outlet user"
			return c.JSON(http.StatusInternalServerError, response)
		} else {

			products := requestBody["products"].([]interface{})

			for i := 0; i < len(products); i++ {
				product := products[i].(map[string]interface{})

				getProduct, err := model.GetOneProductById(fmt.Sprint(product["id"]))

				if err != nil {
					response.Message = "Data tidak ditemukan"
					return c.JSON(http.StatusNotFound, response)
				} else {

					quantity, err := strconv.Atoi(fmt.Sprint(product["quantity"]))
					if err != nil {
						response.Message = "Convert quantity gagal"
						return c.JSON(http.StatusInternalServerError, response)
					}

					transactionDetails := new(structs.TransactionDetails)
					transactionDetails.TransactionId = transaction.Id
					transactionDetails.ProductId = getProduct.Id
					transactionDetails.Quantity = quantity
					transactionDetails.Price = quantity * getProduct.Price
					transactionDetails.Note = fmt.Sprint(product["note"])

					if model.CreateTransactionDetail(transactionDetails) != nil { // method create outlet
						response.Message = "Gagal create transaction detail"
						return c.JSON(http.StatusInternalServerError, response)
					}
				}
			}

			response.Message = "Berhasil buat transaksi"
			return c.JSON(http.StatusOK, response)
		}
	}
}

func TransactionShow(c echo.Context) error {
	transaction, err := model.GetOneTransactionById(c.Param("id")) // method get by email
	response := new(structs.Response)

	if err != nil {
		response.Message = "Data tidak ditemukan"
		return c.JSON(http.StatusNotFound, response)
	} else {
		response.Message = "Sukses melihat data"
		response.Data = transaction
		return c.JSON(http.StatusOK, response)
	}
}

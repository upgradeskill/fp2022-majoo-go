package models

import (
	database "mini-pos/configs"
	"mini-pos/structs"
)

func CreateTransactionDetail(transactionDetail *structs.TransactionDetails) error {
	if err := database.DB.Create(transactionDetail).Error; err != nil {
		return err
	}
	return nil
}

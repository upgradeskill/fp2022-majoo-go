package models

import (
	database "mini-pos/configs"
	"mini-pos/structs"
)

func CreateTransaction(transaction *structs.Transactions) error {
	if err := database.DB.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTransaction(id string, transaction *structs.Transactions) error {
	if err := database.DB.Model(transaction).Where("id = ?", id).Updates(transaction).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTransaction(transaction *structs.Transactions) error {
	if err := database.DB.Delete(transaction).Error; err != nil {
		return err
	}
	return nil
}

func GetOneTransactionById(id string) (structs.Transactions, error) {
	var transaction structs.Transactions
	result := database.DB.
		Preload("Outlet").
		Preload("TransactionDetails.Product.Category").
		Where("id = ?", id).First(&transaction)
	return transaction, result.Error
}

func GetAllTransaction(q string, limit int, offset int) ([]structs.Transactions, error) {
	var transactions []structs.Transactions

	result := database.DB.
		Preload("Outlet").
		Preload("TransactionDetails.Product.Category").
		Where("code LIKE ?", "%"+q+"%").
		Limit(limit).
		Offset(offset).
		Find(&transactions)

	return transactions, result.Error
}

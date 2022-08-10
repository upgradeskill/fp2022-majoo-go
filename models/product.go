package models

import (
	database "mini-pos/configs"
	"mini-pos/structs"
)

func CreateProduct(product *structs.Products) error {
	if err := database.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(id string, product *structs.Products) error {
	if err := database.DB.Model(product).Where("id = ?", id).Updates(product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(product *structs.Products) error {
	if err := database.DB.Delete(product).Error; err != nil {
		return err
	}
	return nil
}

func GetOneProductById(id string) (structs.Products, error) {
	var product structs.Products
	result := database.DB.Where("id = ?", id).First(&product)
	return product, result.Error
}

func GetAllProduct(q string, limit int, offset int) ([]structs.Products, error) {
	var products []structs.Products
	result := database.DB.Where("name LIKE ?", "%"+q+"%").Limit(limit).Offset(offset).Find(&products)

	return products, result.Error
}

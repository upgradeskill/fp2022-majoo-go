package models

import (
	database "mini-pos/configs"
	"mini-pos/structs"
)

func CreateCategory(category *structs.Categories) error {
	if err := database.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCategory(id string, category *structs.Categories) error {
	if err := database.DB.Model(category).Where("id = ?", id).Updates(category).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCategory(category *structs.Categories) error {
	if err := database.DB.Delete(category).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCategoryById(id string) (structs.Categories, error) {
	var category structs.Categories
	result := database.DB.Where("id = ?", id).First(&category)
	return category, result.Error
}

func GetAllCategory(q string) ([]structs.Categories, error) {
	var categories []structs.Categories
	result := database.DB.Where("name LIKE ?", "%"+q+"%").Find(&categories)

	return categories, result.Error
}

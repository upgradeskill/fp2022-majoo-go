package models

import (
	config "mini-pos/configs"
)

type Categories struct {
	Id   int    `json:"id" form:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name"`
}

func (category *Categories) CreateCategory() error {
	if err := config.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (category *Categories) UpdateCategory(id string) error {
	if err := config.DB.Model(&Categories{}).Where("id = ?", id).Updates(category).Error; err != nil {
		return err
	}
	return nil
}

func (category *Categories) DeleteCategory() error {
	if err := config.DB.Delete(category).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCategoryById(id string) (Categories, error) {
	var cateogry Categories
	result := config.DB.Where("id = ?", id).First(&cateogry)
	return cateogry, result.Error
}

func GetAllCategory(keywords string) ([]Categories, error) {
	var cateogries []Categories
	result := config.DB.Where("name LIKE ?", "%"+keywords+"%").Find(&cateogries)

	return cateogries, result.Error
}

package models

import (
	database "mini-pos/configs"
	"mini-pos/structs"
)

func CreateOutlet(outlet *structs.Outlets) error {
	if err := database.DB.Create(outlet).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOutlet(id string, outlet *structs.Outlets) error {
	if err := database.DB.Model(outlet).Where("id = ?", id).Updates(outlet).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOutlet(outlet *structs.Outlets) error {
	if err := database.DB.Delete(outlet).Error; err != nil {
		return err
	}
	return nil
}

func GetOneOutletById(id string) (structs.Outlets, error) {
	var outlet structs.Outlets
	result := database.DB.Where("id = ?", id).First(&outlet)
	return outlet, result.Error
}

func GetAllOutlet(q string) ([]structs.Outlets, error) {
	var outlets []structs.Outlets
	result := database.DB.Where("name LIKE ?", "%"+q+"%").Find(&outlets)

	return outlets, result.Error
}

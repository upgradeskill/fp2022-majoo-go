package models

import (
	database "mini-pos/configs"
	"mini-pos/structs"
)

func CreateUser(user *structs.Users) error {
	if err := database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(id string, user *structs.Users) error {
	if err := database.DB.Model(user).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *structs.Users) error {
	if err := database.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUserById(id string) (structs.Users, error) {
	var user structs.Users
	result := database.DB.Where("id = ?", id).First(&user)
	return user, result.Error
}

func GetOneUserByEmail(email string) (structs.Users, error) {
	var user structs.Users
	result := database.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func GetAllUser(q string, limit int, offset int) ([]structs.Users, error) {
	var users []structs.Users
	result := database.DB.
		Preload("OutletUsers").
		Where("email LIKE ? OR name LIKE ?", "%"+q+"%", "%"+q+"%").Find(&users)

	return users, result.Error
}

package models

import (
	"fmt"
	config "mini-pos/configs"
)

type Users struct {
	Id       int    `json:"id" form:"id" gorm:"primaryKey"`
	Email    string `json:"email" form:"email" `
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	IsAdmin  int    `json:"is_admin" form:"is_admin"`
}

func (user *Users) CreateUser() error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) UpdateUser(id string) error {
	if err := config.DB.Model(&Users{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) DeleteUser() error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func GetOneById(id string) (Users, error) {
	var user Users
	result := config.DB.Where("id = ?", id).First(&user)
	return user, result.Error
}

func GetOneByEmail(email string) (Users, error) {
	fmt.Println("emails", email)
	var user Users
	result := config.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func GetAll(keywords string) ([]Users, error) {
	var users []Users
	result := config.DB.Where("email LIKE ? OR name LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&users)

	return users, result.Error
}

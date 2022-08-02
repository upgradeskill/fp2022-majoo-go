package models

import (
	database "mini-pos/configs"
	"mini-pos/structs"
)

func CreateOutletUser(outletUser *structs.OutletUsers) error {
	if err := database.DB.Create(outletUser).Error; err != nil {
		return err
	}
	return nil
}

func GetOutletUserByUserId(userId string) ([]structs.OutletUsers, error) {
	var outletUsers []structs.OutletUsers
	result := database.DB.Where("user_id = ?", userId).Find(&outletUsers)
	return outletUsers, result.Error
}

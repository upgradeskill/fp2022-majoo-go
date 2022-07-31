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

package entity

type OutletUser struct {
	Id       uint `json:"id" form:"id" gorm:"primaryKey"`
	OutletId uint `json:"outlet_id" form:"outlet_id"`
	UserId   uint `json:"user_id" form:"user_id"`
}

package structs

type OutletUsers struct {
	Id       int `json:"id" form:"id" gorm:"primaryKey"`
	OutletId int `json:"outlet_id" form:"outlet_id"`
	UserId   int `json:"user_id" form:"user_id"`
}

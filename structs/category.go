package structs

type Categories struct {
	Id       int    `json:"id" form:"id" gorm:"primaryKey"`
	OutletId int    `json:"outlet_id" form:"outlet_id"`
	Name     string `json:"name" form:"name"`
}

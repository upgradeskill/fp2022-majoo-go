package entity

type Outlet struct {
	Id          uint         `json:"id" form:"id" gorm:"primaryKey"`
	Name        string       `json:"name" form:"name"`
	OutletUsers []OutletUser `json:"outlet_users" gorm:"Foreignkey:outlet_id;association_foreignkey:Id;"`
}

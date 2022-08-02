package structs

type Outlets struct {
	Id          int           `json:"id" form:"id" gorm:"primaryKey"`
	Name        string        `json:"name" form:"name"`
	OutletUsers []OutletUsers `gorm:"Foreignkey:outlet_id;association_foreignkey:Id;" json:"outlet_users"`
}

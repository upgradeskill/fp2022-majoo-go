package structs

import "time"

type Products struct {
	Id          int        `json:"id" form:"id" gorm:"primaryKey"`
	CategoryId  int        `json:"category_id" form:"category_id"`
	OutletId    int        `json:"outlet_id" form:"outlet_id"`
	Name        string     `json:"name" form:"name"`
	Description string     `json:"description" form:"description"`
	Price       int        `json:"price" form:"price"`
	CreatedAt   time.Time  `json:"created_at" form:"created_at"`
	Category    Categories `gorm:"Foreignkey:category_id;" json:"category"`
}

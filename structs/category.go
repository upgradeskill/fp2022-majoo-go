package structs

type Categories struct {
	Id   int    `json:"id" form:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name"`
}

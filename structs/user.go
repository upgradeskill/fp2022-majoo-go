package structs

type Users struct {
	Id       int    `json:"id" form:"id" gorm:"primaryKey"`
	Email    string `json:"email" form:"email" `
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	IsAdmin  int    `json:"is_admin" form:"is_admin"`
}

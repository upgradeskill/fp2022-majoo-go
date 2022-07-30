package structs

type Users struct {
	Email string `json:"email" form:"email" gorm:"primaryKey" binding:"required"`
	Name  string `json:"name" form:"name" binding:"required"`
}

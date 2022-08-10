package entity

import "time"

type User struct {
	Id           uint      `json:"id" form:"id" gorm:"primaryKey"`
	Email        string    `json:"email" form:"email" `
	Name         string    `json:"name" form:"name"`
	Password     string    `json:"password" form:"password"`
	PasswordSalt string    `json:"password_salt" form:"password_salt"`
	IsAdmin      bool      `json:"is_admin" form:"is_admin"`
	Created      time.Time `json:"created"`
	Modified     time.Time `json:"modified"`
}

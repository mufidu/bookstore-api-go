package models

import "time"

type LoginUser struct {
	Username string `json:"username" binding:"required" example:"test"`
	Password string `json:"password" binding:"required" example:"Test123!"`
}

type RegisterUser struct {
	Username string `json:"username" binding:"required" example:"test"`
	Password string `json:"password" binding:"required" example:"Test123!"`
	Email    string `json:"email" binding:"required" example:"test@test.com"`
}

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

package models

import "time"

type LoginAdmin struct {
    Username string `json:"username" example:"admin" binding:"required"`
    Password string `json:"password" example:"admin123!" binding:"required"`
}

type RegisterAdmin struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"admin123!"`
	Email    string `json:"email" binding:"required" example:"admin@test.com"`
}

type Admin struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

package models

import "time"

type Cart struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	TotalPrice int       `json:"total_price"`
	UserID     uint      `json:"user_id"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	Books      []*Book   `json:"books" gorm:"many2many:cart_books;"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateCart struct {
	TotalPrice int `json:"total_price" binding:"required"`
	UserID     int `json:"user_id" binding:"required"`
}

type UpdateCart struct {
	TotalPrice int `json:"total_price"`
	UserID     int `json:"user_id"`
}

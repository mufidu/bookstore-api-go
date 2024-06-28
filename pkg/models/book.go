package models

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	Genre     string    `json:"genre"`
	Cover     string    `json:"cover"`
	Author    string    `json:"author"`
	Year      int       `json:"year"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateBook struct {
	Title    string `json:"title" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	Genre    string `json:"genre" binding:"required"`
	Cover    string `json:"cover" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Year     int    `json:"year" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type UpdateBook struct {
	Title    string `json:"title"`
	Price    int    `json:"price"`
	Genre    string `json:"genre"`
	Cover    string `json:"cover"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Quantity int    `json:"quantity"`
}

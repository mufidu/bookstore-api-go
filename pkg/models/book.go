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
	Carts     []*Cart   `json:"carts" gorm:"many2many:cart_books;"`
	Users     []*User   `json:"users" gorm:"many2many:user_books;"`
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

type CartBook struct {
	BookID   uint `json:"book_id"`
	CartID   uint `json:"cart_id"`
	Quantity int  `json:"quantity"`
}

type UserBook struct {
	BookID   uint `json:"book_id"`
	UserID   uint `json:"user_id"`
	Quantity int  `json:"quantity"`
}

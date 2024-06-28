package models

import (
	"time"
)

type Transaction struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	InvoiceNumber string    `json:"invoice_number"`
	Status        string    `json:"status"`
	Amount        int       `json:"amount"`
	Items         string    `json:"items"`
	QrisString    string    `json:"qris_string"`
	QrisURL       string    `json:"qris_url"`
	ExpiryTime    string    `json:"expiry_time"`
	InvoiceDate   string    `json:"invoice_date"`
	UserID        uint      `json:"user_id"`
	User          User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateTransaction struct {
	InvoiceNumber string `json:"invoice_number" binding:"required"`
	Status        string `json:"status" binding:"required"`
	Amount        int    `json:"amount" binding:"required"`
	Items         string `json:"items" binding:"required"`
	QrisString    string `json:"qris_string" binding:"required"`
	QrisURL       string `json:"qris_url" binding:"required"`
	ExpiryTime    string `json:"expiry_time" binding:"required"`
	InvoiceDate   string `json:"invoice_date" binding:"required"`
	UserID        int    `json:"user_id" binding:"required"`
}

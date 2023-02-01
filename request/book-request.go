package request

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      int         `json:"rating" binding:"required,number"`
	Discount    int         `json:"discount" binding:"required,number"`
	Quantity    json.Number `json:"quantity" binding:"required,number"`
	GenreID     json.Number `json:"genre_id" binding:"required,number"`
	StatusID    json.Number `json:"status_id" binding:"required,number"`
}

type BuyBookRequest struct {
	Quantity           int    `json:"quantity" binding:"required,number"`
	PaymentBankID      int    `json:"payment_bank_id" binding:"required,number"`
	PaymentBankAccount string `json:"payment_bank_account" binding:"required"`
}

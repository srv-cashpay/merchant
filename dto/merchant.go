package dto

import (
	"time"

	"gorm.io/gorm"
)

type GetMerchantByIdRequest struct {
	ID string `query:"id" validate:"required, id"`
}

type GetMerchantRequest struct {
	ID           string         `json:"id"`
	UserID       string         `json:"user_id"`
	IDNumber     string         `json:"id_number"`
	MerchantID   string         `json:"merchant_id"`
	MerchantName string         `json:"merchant_name"`
	Address      string         `json:"address"`
	Country      string         `json:"country"`
	City         string         `json:"city"`
	Zip          string         `json:"zip"`
	Phone        string         `json:"phone"`
	CurrencyID   string         `json:"currency_id"`
	Description  string         `json:"description"`
	UpdatedBy    string         `json:"update_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type GetMerchantResponse struct {
	ID           string         `json:"id"`
	UserID       string         `json:"user_id"`
	MerchantID   string         `json:"merchant_id"`
	IDNumber     string         `json:"id_number"`
	MerchantName string         `json:"merchant_name"`
	Address      string         `json:"address"`
	Country      string         `json:"country"`
	City         string         `json:"city"`
	Zip          string         `json:"zip"`
	Phone        string         `json:"phone"`
	CurrencyID   int            `json:"currency_id"`
	Description  string         `json:"description"`
	UpdatedBy    string         `json:"update_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type UpdateMerchantRequest struct {
	ID           string    `query:"id" validate:"required, id"`
	UserID       string    `json:"user_id"`
	MerchantName string    `json:"merchant_name"`
	IDNumber     string    `json:"id_number"`
	Address      string    `json:"address"`
	Country      string    `json:"country"`
	City         string    `json:"city"`
	Zip          string    `json:"zip"`
	Phone        string    `json:"phone"`
	CurrencyID   int       `json:"currency_id"`
	Description  string    `json:"description"`
	UpdatedBy    string    `json:"update_by"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateMerchantResponse struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	MerchantName string    `json:"merchant_name"`
	IDNumber     string    `json:"id_number"`
	Address      string    `json:"address"`
	Country      string    `json:"country"`
	City         string    `json:"city"`
	Zip          string    `json:"zip"`
	Phone        string    `json:"phone"`
	CurrencyID   int       `json:"currency_id"`
	Description  string    `json:"description"`
	UpdatedBy    string    `json:"update_by"`
	UpdatedAt    time.Time `json:"updated_at"`
}

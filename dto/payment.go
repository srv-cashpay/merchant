package dto

import (
	"time"
)

type PaymentRequest struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	MerchantID        string    `json:"merchant_id"`
	PaymentName       string    `json:"payment_name"`
	PaymentPercentage uint      `json:"payment_percentage"`
	Status            int       `json:"status"`
	Description       string    `json:"description"`
	CreatedBy         string    `json:"created_by"`
	UpdatedBy         string    `json:"updated_by"`
	DeletedBy         string    `json:"deleted_by"`
	CreatedAt         time.Time `json:"created_at"`
}

type PaymentResponse struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	MerchantID        string    `json:"merchant_id"`
	PaymentName       string    `json:"payment_name"`
	PaymentPercentage uint      `json:"payment_percentage"`
	Status            string    `json:"status"`
	Description       string    `json:"description"`
	CreatedBy         string    `json:"created_by"`
	UpdatedBy         string    `json:"updated_by"`
	DeletedBy         string    `json:"deleted_by"`
	CreatedAt         time.Time `json:"created_at"`
}
type PaymentUpdateRequest struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	MerchantID        string    `json:"merchant_id"`
	PaymentName       string    `json:"payment_name"`
	PaymentPercentage uint      `json:"payment_percentage"`
	Status            int       `json:"status"`
	Description       string    `json:"description"`
	CreatedBy         string    `json:"created_by"`
	UpdatedBy         string    `json:"updated_by"`
	DeletedBy         string    `json:"deleted_by"`
	CreatedAt         time.Time `json:"created_at"`
}

type PaymentUpdateResponse struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	MerchantID        string    `json:"merchant_id"`
	PaymentName       string    `json:"payment_name"`
	PaymentPercentage uint      `json:"payment_percentage"`
	Status            int       `json:"status"`
	Description       string    `json:"description"`
	CreatedBy         string    `json:"created_by"`
	UpdatedBy         string    `json:"updated_by"`
	DeletedBy         string    `json:"deleted_by"`
	CreatedAt         time.Time `json:"created_at"`
}

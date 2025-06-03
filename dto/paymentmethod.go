package dto

import (
	"time"
)

type PaymentMethodRequest struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantID    string    `json:"merchant_id"`
	PaymentMethod string    `json:"payment_method"`
	Status        int       `json:"status"`
	FileName      string    `json:"file_name"`
	FilePath      string    `json:"file_path"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type PaymentMethodResponse struct {
	ID            uint      `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantID    string    `json:"merchant_id"`
	PaymentMethod string    `json:"payment_method"`
	Status        int       `json:"status"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
}
type PaymentMethodUpdateRequest struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantID    string    `json:"merchant_id"`
	PaymentMethod string    `json:"payment_method"`
	Status        int       `json:"status"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type PaymentMethodUpdateResponse struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantID    string    `json:"merchant_id"`
	PaymentMethod string    `json:"payment_method"`
	Status        int       `json:"status"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type GetByIdPaymentRequest struct {
	ID uint `param:"id" validate:"required"`
}

type DeletePaymentRequest struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeletePaymentResponse struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

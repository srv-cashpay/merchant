package dto

import (
	"time"
)

type UserRequest struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	User        string    `json:"user"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	User        string    `json:"user"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}
type UserUpdateRequest struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	User        string    `json:"user"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserUpdateResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	User        string    `json:"user"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

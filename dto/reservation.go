package dto

import (
	"time"
)

type ReservationRequest struct {
	ID          uint      `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	Name        string    `json:"name"`
	Whatsapp    string    `json:"whatsapp"`
	Date        string    `json:"date"`
	Time        string    `json:"time"`
	Table       []Table   `json:"table"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type ReservationResponse struct {
	ID          uint      `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	Table       []Table   `json:"table"`
	Name        string    `json:"name"`
	Whatsapp    string    `json:"whatsapp"`
	Date        string    `json:"date"`
	Time        string    `json:"time"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateReservationRequest struct {
	ID          uint      `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	Name        string    `json:"name"`
	Whatsapp    string    `json:"whatsapp"`
	Date        string    `json:"date"`
	Time        string    `json:"time"`
	Table       []Table   `json:"table"`
	Floor       string    `json:"floor"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateReservationResponse struct {
	ID          uint      `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	Name        string    `json:"name"`
	Whatsapp    string    `json:"whatsapp"`
	Date        string    `json:"date"`
	Time        string    `json:"time"`
	Table       []Table   `json:"table"`
	Floor       string    `json:"floor"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type Table struct {
	Table string `json:"table"`
	Floor string `json:"floor"`
}

type DeleteReservationRequest struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeleteReservationResponse struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type GetReservationByIdRequest struct {
	ID uint `param:"id" validate:"required"`
}

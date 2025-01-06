package dto

import (
	"time"
)

type ReservationRequest struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Reservation string `json:"resevation"`
	Name        string
	Whatsapp    string
	Date        string
	Time        string
	Table       []TableResponse
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type ReservationResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Reservation string    `json:"resevation"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

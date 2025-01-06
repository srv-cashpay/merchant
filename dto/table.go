package dto

import (
	"time"
)

type TableRequest struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Table       string    `json:"resevation"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type TableResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Table       string    `json:"resevation"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}
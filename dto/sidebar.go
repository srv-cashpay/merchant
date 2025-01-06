package dto

import (
	"time"
)

type SidebarRequest struct {
	ID         uint      `json:"id"`
	UserID     string    `json:"user_id"`
	MerchantID string    `json:"merchant_id"`
	Label      string    `json:"label"`
	Icon       string    `json:"icon"`
	To         string    `json:"to"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedBy  string    `json:"deleted_by"`
	CreatedAt  time.Time `json:"created_at"`
}

type SidebarResponse struct {
	ID         uint      `json:"id"`
	UserID     string    `json:"user_id"`
	MerchantID string    `json:"merchant_id"`
	Label      string    `json:"label"`
	Icon       string    `json:"icon"`
	To         string    `json:"to"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedBy  string    `json:"deleted_by"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetSidebarByIdRequest struct {
	ID         uint      `json:"id"`
	UserID     string    `json:"user_id"`
	MerchantID string    `json:"merchant_id"`
	Label      string    `json:"label"`
	Icon       string    `json:"icon"`
	To         string    `json:"to"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedBy  string    `json:"deleted_by"`
	CreatedAt  time.Time `json:"created_at"`
}

type DeleteSidebarRequest struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeleteSidebarResponse struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type GetSidebarResponse struct {
	Items []SidebarItem `json:"items"`
}

type SidebarItem struct {
	Label string `json:"label"`
	Icon  string `json:"icon"`
	To    string `json:"to"`
}

type SidebarUpdateRequest struct {
	ID         string `json:"id"`
	Label      string `json:"label"`
	Icon       string `json:"icon"`
	To         string `json:"to"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	UpdatedBy  string `json:"updated_by"`
}

type SidebarUpdateResponse struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	MerchantID string    `json:"merchant_id"`
	Label      string    `json:"label"`
	Icon       string    `json:"icon"`
	To         string    `json:"to"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedBy  string    `json:"deleted_by"`
	CreatedAt  time.Time `json:"created_at"`
}

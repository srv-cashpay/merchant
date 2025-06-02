package dto

import "time"

type RoleRequest struct {
	ID         string `json:"id"`
	Role       string `json:"role_id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
}

type RoleResponse struct {
	ID         string `json:"id"`
	Role       string `json:"role_id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
}

type RoleUpdateRequest struct {
	ID         string `json:"id"`
	Role       string `json:"role_id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
	UpdatedBy  string `json:"updated_by"`
}

type RoleUpdateResponse struct {
	ID         string `json:"id"`
	Role       string `json:"role_id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
	UpdatedBy  string `json:"updated_by"`
}

type DeleteRoleRequest struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeleteRoleResponse struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type GetRoleResponse struct {
	Items []RoleResponse `json:"items"`
}

type GetRoleByIdRequest struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	MerchantID string    `json:"merchant_id"`
	Role       string    `json:"role_id"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedBy  string    `json:"deleted_by"`
	CreatedAt  time.Time `json:"created_at"`
}

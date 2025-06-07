package dto

import "time"

type RoleUserRequest struct {
	ID         string `json:"id"`
	Role       string `json:"role_id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
}

type RoleUserResponse struct {
	ID         string `json:"id"`
	Role       string `json:"role`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
}

type RoleUserUpdateRequest struct {
	ID         string `json:"id"`
	Role       string `json:"role"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
	UpdatedBy  string `json:"updated_by"`
}

type RoleUserUpdateResponse struct {
	ID         string `json:"id"`
	Role       string `json:"role"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
	UpdatedBy  string `json:"updated_by"`
}

type DeleteRoleUserRequest struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeleteRoleUserResponse struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type GetRoleUserResponse struct {
	Items []RoleUserResponse `json:"items"`
}

type GetRoleUserByIdRequest struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	MerchantID string    `json:"merchant_id"`
	Role       string    `json:"role"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedBy  string    `json:"deleted_by"`
	CreatedAt  time.Time `json:"created_at"`
}

package dto

import "time"

type RoleUserRequest struct {
	ID           uint   `json:"id"`
	RoleID       string `json:"role_id"`
	PermissionID uint   `json:"permission_id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	CreatedBy    string `json:"created_by"`
}

type RoleUserResponse struct {
	ID           uint   `json:"id"`
	RoleID       string `json:"role_id"`
	PermissionID uint   `json:"permission_id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	CreatedBy    string `json:"created_by"`
}

type RoleUserUpdateRequest struct {
	ID           uint   `json:"id"`
	RoleID       string `json:"role_id"`
	PermissionID uint   `json:"permission_id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	CreatedBy    string `json:"created_by"`
	UpdatedBy    string `json:"updated_by"`
}

type RoleUserUpdateResponse struct {
	ID           uint   `json:"id"`
	RoleID       string `json:"role_id"`
	PermissionID uint   `json:"permission_id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	CreatedBy    string `json:"created_by"`
	UpdatedBy    string `json:"updated_by"`
}

type DeleteRoleUserRequest struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeleteRoleUserResponse struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type GetRoleUserResponse struct {
	Items []RoleUserResponse `json:"items"`
}

type GetRoleUserByIdRequest struct {
	ID           uint      `json:"id"`
	UserID       string    `json:"user_id"`
	MerchantID   string    `json:"merchant_id"`
	RoleID       string    `json:"role_id"`
	PermissionID uint      `json:"permission_id"`
	CreatedBy    string    `json:"created_by"`
	UpdatedBy    string    `json:"updated_by"`
	DeletedBy    string    `json:"deleted_by"`
	CreatedAt    time.Time `json:"created_at"`
}

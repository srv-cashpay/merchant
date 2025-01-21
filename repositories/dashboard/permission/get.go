package permission

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (r *PermissionRepository) Get(req dto.RoleUser) (dto.GetPermissionResponse, error) {
	var permissions []dto.PermissionResponse

	// Query with JOIN to fetch permissions
	err := r.DB.Table("role_users").
		Select("permissions.label, permissions.icon, permissions.to").
		Joins("JOIN permissions ON role_users.permission_id = permissions.id").
		Where("role_users.user_id = ?", req.UserID).
		Scan(&permissions).Error

	if err != nil {
		return dto.GetPermissionResponse{}, err
	}

	// Wrap the permissions in the 'items' field as expected
	return dto.GetPermissionResponse{
		Items: permissions,
	}, nil
}

package roleuser

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (r *RoleUserRepository) Get(req dto.RoleUserRequest) (dto.GetRoleUserResponse, error) {
	var roles []dto.RoleUserResponse

	err := r.DB.Table("role_users").
		Select("roles.role AS label, roles.id AS role_id, roles.merchant_id, roles.user_id, roles.created_at").
		Joins("JOIN roles ON role_users.role_id = roles.id").
		Where("role_users.user_id = ?", req.UserID).
		Scan(&roles).Error

	if err != nil {
		return dto.GetRoleUserResponse{}, err
	}

	return dto.GetRoleUserResponse{
		Items: roles,
	}, nil
}

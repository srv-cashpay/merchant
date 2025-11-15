package roleuser

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (r *RoleUserRepository) Get(req dto.RoleUserRequest) (dto.GetRoleUserResponse, error) {
	var roles []dto.RoleUserResponse

	err := r.DB.Table("role_users AS ru").
		Select(`
            roles.role AS role_id,       
            roles.role AS label,        
            roles.merchant_id,
            roles.user_id,
            roles.created_at
        `).
		Joins("JOIN roles ON roles.id = ru.role_id").
		Where("ru.user_id = ?", req.UserID).
		Scan(&roles).Error

	if err != nil {
		return dto.GetRoleUserResponse{}, err
	}

	return dto.GetRoleUserResponse{
		Items: roles,
	}, nil
}

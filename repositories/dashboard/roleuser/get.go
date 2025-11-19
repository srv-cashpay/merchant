package roleuser

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (r *RoleUserRepository) Get(req dto.RoleUserRequest) (dto.GetRoleUserResponse, error) {
	var result []dto.RoleUserResponse

	err := r.DB.Table("role_users AS ru").
		Select(`
			ru.id,
			ru.role_id,
			ru.permission_id,
			ru.user_id,
			ru.merchant_id,
			p.created_by
		`).
		Joins("LEFT JOIN permissions AS p ON JSON_CONTAINS(ru.permission_id, JSON_QUOTE(CAST(p.id AS CHAR)))").
		Where("ru.user_id = ? AND ru.merchant_id = ?", req.UserID, req.MerchantID).
		Scan(&result).Error

	if err != nil {
		return dto.GetRoleUserResponse{}, err
	}

	// Bungkus izin dalam field 'items' seperti yang diharapkan
	return dto.GetRoleUserResponse{
		Items: result,
	}, nil
}

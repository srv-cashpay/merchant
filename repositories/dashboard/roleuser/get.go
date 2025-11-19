package roleuser

import (
	"encoding/json"

	dto "github.com/srv-cashpay/merchant/dto"
)

func (r *RoleUserRepository) Get(req dto.RoleUserRequest) (dto.GetRoleUserResponse, error) {
	var rows []struct {
		ID           uint
		RoleID       string
		PermissionID []byte
		UserID       string
		MerchantID   string
		CreatedBy    string
	}

	err := r.DB.Table("role_users AS ru").
		Select(`
			ru.id,
			ru.role_id,
			ru.permission_id,
			ru.user_id,
			ru.merchant_id,
			p.created_by
		`).
		Joins(`
			LEFT JOIN permissions AS p 
				ON ru.permission_id::jsonb @> ('[' || p.id || ']')::jsonb
		`).
		Where("ru.user_id = ?", req.UserID).
		Where("ru.merchant_id = ?", req.MerchantID).
		Scan(&rows).Error

	if err != nil {
		return dto.GetRoleUserResponse{}, err
	}

	var result []dto.RoleUserResponse

	for _, row := range rows {
		var permIDs []uint
		json.Unmarshal(row.PermissionID, &permIDs)

		result = append(result, dto.RoleUserResponse{
			ID:           row.ID,
			RoleID:       row.RoleID,
			PermissionID: permIDs,
			UserID:       row.UserID,
			MerchantID:   row.MerchantID,
			CreatedBy:    row.CreatedBy,
		})
	}

	return dto.GetRoleUserResponse{
		Items: result,
	}, nil
}

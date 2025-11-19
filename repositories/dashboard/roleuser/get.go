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

		PermID uint
		Label  string
		Icon   string
		To     string
	}

	err := r.DB.Table("role_users AS ru").
		Select(`
			ru.id,
			ru.role_id,
			ru.permission_id,
			ru.user_id,
			ru.merchant_id,
			p.id AS perm_id,
			p.label,
			p.icon,
			p.to
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

	// Map ke response
	roleMap := map[uint]*dto.RoleUserResponse{}

	for _, row := range rows {
		// jika belum ada -> buat object baru
		if _, ok := roleMap[row.ID]; !ok {

			var permIDs []uint
			json.Unmarshal(row.PermissionID, &permIDs)

			roleMap[row.ID] = &dto.RoleUserResponse{
				ID:           row.ID,
				RoleID:       row.RoleID,
				PermissionID: permIDs,
				UserID:       row.UserID,
				MerchantID:   row.MerchantID,
				Permissions:  []dto.PermissionItem{},
			}
		}

		// Append permission jika ada
		if row.PermID != 0 {
			roleMap[row.ID].Permissions = append(roleMap[row.ID].Permissions, dto.PermissionItem{
				ID:    row.PermID,
				Label: row.Label,
				Icon:  row.Icon,
				To:    row.To,
			})
		}
	}

	// Convert map → array
	var result []dto.RoleUserResponse
	for _, v := range roleMap {
		result = append(result, *v)
	}

	return dto.GetRoleUserResponse{
		Roles: result,
	}, nil
}

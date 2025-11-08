package role

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *RoleRepository) RoleUser(req dto.GetRoleRequest) (dto.GetRoleResponse, error) {
	var roles []entity.Role

	if err := r.DB.
		Find(&roles).Error; err != nil {
		return dto.GetRoleResponse{}, err
	}

	items := make([]dto.RoleResponse, len(roles))
	for i, role := range roles {
		items[i] = dto.RoleResponse{
			Role: role.Role,
		}
	}

	return dto.GetRoleResponse{Items: items}, nil
}

package roleuserpermission

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *RoleUserPermissionRepository) Create(req dto.RoleUserPermissionRequest) (dto.RoleUserPermissionResponse, error) {

	// Create the new Role entry
	create := entity.RoleUserPermission{
		ID: req.ID,
	}

	// Save the new Role to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.RoleUserPermissionResponse{}, err
	}

	// Build the response for the created Role
	response := dto.RoleUserPermissionResponse{
		ID: create.ID,
	}

	return response, nil
}

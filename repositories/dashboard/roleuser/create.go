package roleuser

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *RoleUserRepository) Create(req dto.RoleUserRequest) (dto.RoleUserResponse, error) {
	// Create the new Role entry
	create := entity.RoleUser{
		RoleID:       req.RoleID,
		UserID:       req.UserID,
		PermissionID: req.PermissionID,
	}

	// Save the new Role to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.RoleUserResponse{}, err
	}

	// Build the response for the created Role
	response := dto.RoleUserResponse{
		RoleID:       create.RoleID,
		UserID:       req.UserID,
		PermissionID: req.PermissionID,
	}

	return response, nil
}

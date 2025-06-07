package roleuserpermission

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *roleuserpermissionService) Create(req dto.RoleUserPermissionRequest) (dto.RoleUserPermissionResponse, error) {

	create := dto.RoleUserPermissionRequest{
		RoleUserID: req.RoleUserID,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.RoleUserPermissionResponse{}, err
	}

	response := dto.RoleUserPermissionResponse{
		RoleUserID: created.RoleUserID,
	}

	return response, nil
}

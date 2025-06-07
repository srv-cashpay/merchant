package roleuser

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *roleuserService) Create(req dto.RoleUserRequest) (dto.RoleUserResponse, error) {

	create := dto.RoleUserRequest{
		Role:       req.Role,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.RoleUserResponse{}, err
	}

	response := dto.RoleUserResponse{
		Role: created.Role,
	}

	return response, nil
}

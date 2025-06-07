package roleuser

import "github.com/srv-cashpay/merchant/dto"

func (b *roleuserService) Update(req dto.RoleUserUpdateRequest) (dto.RoleUserUpdateResponse, error) {
	request := dto.RoleUpdateRequest{
		Role:      req.Role,
		UpdatedBy: req.UpdatedBy,
		UserID:    req.UserID,
	}

	role, err := b.Repo.Update(req)
	if err != nil {
		return role, err
	}

	response := dto.RoleUserUpdateResponse{
		Role:      req.Role,
		UpdatedBy: request.UpdatedBy,
		UserID:    request.UserID,
	}

	return response, nil
}

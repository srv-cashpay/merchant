package user

import "github.com/srv-cashpay/merchant/dto"

func (b *userService) Update(req dto.UserMerchantUpdateRequest) (dto.UserMerchantUpdateResponse, error) {
	request := dto.UserMerchantUpdateRequest{
		FullName:    req.FullName,
		Status:      req.Status,
		UpdatedBy:   req.UpdatedBy,
		Description: req.Description,
	}

	user, err := b.Repo.Update(req)
	if err != nil {
		return user, err
	}

	response := dto.UserMerchantUpdateResponse{
		FullName:    request.FullName,
		Status:      request.Status,
		UpdatedBy:   request.UpdatedBy,
		Description: request.Description,
	}

	return response, nil
}

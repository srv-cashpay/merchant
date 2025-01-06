package user

import "github.com/srv-cashpay/merchant/dto"

func (b *userService) Update(req dto.UserUpdateRequest) (dto.UserUpdateResponse, error) {
	request := dto.UserUpdateRequest{
		User:        req.User,
		Status:      req.Status,
		UpdatedBy:   req.UpdatedBy,
		UserID:      req.UserID,
		Description: req.Description,
	}

	user, err := b.Repo.Update(req)
	if err != nil {
		return user, err
	}

	response := dto.UserUpdateResponse{
		User:        request.User,
		Status:      request.Status,
		UpdatedBy:   request.UpdatedBy,
		UserID:      request.UserID,
		Description: request.Description,
	}

	return response, nil
}

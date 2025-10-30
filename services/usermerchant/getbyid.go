package user

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *userService) GetById(req dto.GetByIdRequest) (*dto.UserMerchantResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

package discount

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *discountService) GetById(req dto.GetByIdRequest) (*dto.DiscountResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

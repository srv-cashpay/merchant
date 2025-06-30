package order

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *orderService) GetById(req dto.GetByIdOrderRequest) (*dto.OrderResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

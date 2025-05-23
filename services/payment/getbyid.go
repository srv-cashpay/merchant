package payment

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *paymentService) GetById(req dto.GetByIdRequest) (*dto.PaymentResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

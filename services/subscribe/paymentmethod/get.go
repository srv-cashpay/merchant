package paymentmethod

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *paymentmethodService) Get(req dto.PaymentMethodRequest) (dto.PaymentMethodResponse, error) {
	pay, err := s.Repo.Get(req)
	if err != nil {
		return dto.PaymentMethodResponse{}, err
	}

	return pay, nil
}

package paymentmethod

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *paymentmethodService) Get(req dto.PaymentMethodRequest) ([]dto.PaymentMethodResponse, error) {
	return s.Repo.Get(req)
}

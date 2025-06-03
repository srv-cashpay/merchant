package paymentmethod

import (
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/subscribe/paymentmethod"
)

type PaymentMethodService interface {
	Create(req dto.PaymentMethodRequest) (dto.PaymentMethodResponse, error)
	Get(req dto.PaymentMethodRequest) (dto.PaymentMethodResponse, error)
	GetById(req dto.GetByIdRequest) (dto.PaymentMethodResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.PaymentMethodUpdateRequest) (dto.PaymentMethodUpdateResponse, error)
}

type paymentmethodService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPaymentMethodService(Repo r.DomainRepository, jwtS m.JWTService) PaymentMethodService {
	return &paymentmethodService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

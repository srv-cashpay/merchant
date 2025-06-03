package paymentmethod

import (
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/subscribe/paymentmethod"
)

type PaymentService interface {
	Create(req dto.PaymentRequest) (dto.PaymentResponse, error)
	Get(req dto.PaymentMethodRequest) (dto.PaymentMethodResponse, error)
	GetById(req dto.GetByIdRequest) (*dto.PaymentResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.PaymentUpdateRequest) (dto.PaymentUpdateResponse, error)
}

type paymentmethodService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPaymentService(Repo r.DomainRepository, jwtS m.JWTService) PaymentService {
	return &paymentmethodService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

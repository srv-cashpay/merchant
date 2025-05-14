package payment

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/payment"
)

type PaymentService interface {
	Create(req dto.PaymentRequest) (dto.PaymentResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdRequest) (*dto.PaymentResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.PaymentUpdateRequest) (dto.PaymentUpdateResponse, error)
}

type paymentService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPaymentService(Repo r.DomainRepository, jwtS m.JWTService) PaymentService {
	return &paymentService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

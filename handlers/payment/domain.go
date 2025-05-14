package payment

import (
	s "github.com/srv-cashpay/merchant/services/payment"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Delete(c echo.Context) error
	BulkDelete(c echo.Context) error
	GetById(c echo.Context) error
	Update(c echo.Context) error
}

type domainHandler struct {
	servicePayment s.PaymentService
}

func NewPaymentHandler(service s.PaymentService) DomainHandler {
	return &domainHandler{
		servicePayment: service,
	}
}

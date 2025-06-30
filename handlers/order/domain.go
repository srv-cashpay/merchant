package order

import (
	s "github.com/srv-cashpay/merchant/services/order"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	GetById(c echo.Context) error
	BulkDelete(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
}

type domainHandler struct {
	serviceOrder s.OrderService
}

func NewOrderHandler(service s.OrderService) DomainHandler {
	return &domainHandler{
		serviceOrder: service,
	}
}

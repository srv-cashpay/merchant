package discount

import (
	s "github.com/srv-cashpay/merchant/services/discount"

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
	serviceDiscount s.DiscountService
}

func NewDiscountHandler(service s.DiscountService) DomainHandler {
	return &domainHandler{
		serviceDiscount: service,
	}
}

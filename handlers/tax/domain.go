package tax

import (
	s "github.com/srv-cashpay/merchant/services/tax"

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
	serviceTax s.TaxService
}

func NewTaxHandler(service s.TaxService) DomainHandler {
	return &domainHandler{
		serviceTax: service,
	}
}

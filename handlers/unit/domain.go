package unit

import (
	s "github.com/srv-cashpay/merchant/services/unit"

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
	serviceUnit s.UnitService
}

func NewUnitHandler(service s.UnitService) DomainHandler {
	return &domainHandler{
		serviceUnit: service,
	}
}

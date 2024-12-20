package pos

import (
	s "github.com/srv-cashpay/merchant/services/pos"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
}

type domainHandler struct {
	servicePos s.PosService
}

func NewPosHandler(service s.PosService) DomainHandler {
	return &domainHandler{
		servicePos: service,
	}
}

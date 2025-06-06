package printer

import (
	s "github.com/srv-cashpay/merchant/services/printer"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type domainHandler struct {
	servicePrinter s.PrinterService
}

func NewPrinterHandler(service s.PrinterService) DomainHandler {
	return &domainHandler{
		servicePrinter: service,
	}
}

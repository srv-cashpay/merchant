package packages

import (
	s "github.com/srv-cashpay/merchant/services/packages"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	MidtransCallback(c echo.Context) error
	ChargeBni(c echo.Context) error
	ChargeBca(c echo.Context) error
	ChargeMandiri(c echo.Context) error
	ChargeBri(c echo.Context) error

	ChargeGpay(c echo.Context) error

	ChargeQris(c echo.Context) error
	ChargeGopay(c echo.Context) error
	ChargeShopeePay(c echo.Context) error
}

type domainHandler struct {
	servicePackages s.PackagesService
}

func NewPackagesHandler(service s.PackagesService) DomainHandler {
	return &domainHandler{
		servicePackages: service,
	}
}

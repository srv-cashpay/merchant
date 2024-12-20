package packages

import (
	s "github.com/srv-cashpay/merchant/services/packages"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	MidtransCallback(c echo.Context) error
}

type domainHandler struct {
	servicePackages s.PackagesService
}

func NewPackagesHandler(service s.PackagesService) DomainHandler {
	return &domainHandler{
		servicePackages: service,
	}
}

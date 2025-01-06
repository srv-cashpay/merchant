package sidebar

import (
	s "github.com/srv-cashpay/merchant/services/dashboard/sidebar"

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
	serviceSidebar s.SidebarService
}

func NewSidebarHandler(service s.SidebarService) DomainHandler {
	return &domainHandler{
		serviceSidebar: service,
	}
}

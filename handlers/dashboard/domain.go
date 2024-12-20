package dashboard

import (
	"github.com/labstack/echo/v4"
	s "github.com/srv-cashpay/merchant/services/dashboard"
)

type DomainHandler interface {
	Get(c echo.Context) error
}

type domainHandler struct {
	serviceDashboard s.DashboardService
}

func NewDashboardHandler(service s.DashboardService) DomainHandler {
	return &domainHandler{
		serviceDashboard: service,
	}
}

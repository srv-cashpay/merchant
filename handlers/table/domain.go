package table

import (
	s "github.com/srv-cashpay/merchant/services/table"

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
	serviceTable s.TableService
}

func NewTableHandler(service s.TableService) DomainHandler {
	return &domainHandler{
		serviceTable: service,
	}
}

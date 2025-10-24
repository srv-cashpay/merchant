package export_data

import (
	s "github.com/srv-cashpay/merchant/services/product/export_data"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	ExportExcel(c echo.Context) error
}

type domainHandler struct {
	serviceExport s.ExportService
}

func NewExportHandler(service s.ExportService) DomainHandler {
	return &domainHandler{
		serviceExport: service,
	}
}

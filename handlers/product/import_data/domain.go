package import_data

import (
	s "github.com/srv-cashpay/merchant/services/product/import_data"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	DownloadTemplate(c echo.Context) error
	UploadProducts(c echo.Context) error
}

type domainHandler struct {
	serviceImport s.ImportService
}

func NewImportHandler(service s.ImportService) DomainHandler {
	return &domainHandler{
		serviceImport: service,
	}
}

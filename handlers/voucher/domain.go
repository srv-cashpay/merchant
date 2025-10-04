package voucher

import (
	s "github.com/srv-cashpay/merchant/services/voucher"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	GetVerifikasi(c echo.Context) error
	GetById(c echo.Context) error
	BulkDelete(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
}

type domainHandler struct {
	serviceVoucher s.VoucherService
}

func NewVoucherHandler(service s.VoucherService) DomainHandler {
	return &domainHandler{
		serviceVoucher: service,
	}
}

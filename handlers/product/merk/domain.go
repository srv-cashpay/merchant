package merk

import (
	"github.com/labstack/echo/v4"
	s "github.com/srv-cashpay/merchant/services/product/merk"
)

type DomainHandler interface {
	Get(c echo.Context) error
}

type domainHandler struct {
	serviceGetMerk s.GetMerkService
}

func NewMerkHandler(service s.GetMerkService) DomainHandler {
	return &domainHandler{
		serviceGetMerk: service,
	}
}

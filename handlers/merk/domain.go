package merk

import (
	s "github.com/srv-cashpay/merchant/services/merk"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	GetById(c echo.Context) error
	BulkDelete(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
}

type domainHandler struct {
	serviceMerk s.MerkService
}

func NewMerkHandler(service s.MerkService) DomainHandler {
	return &domainHandler{
		serviceMerk: service,
	}
}

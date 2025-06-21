package category

import (
	s "github.com/srv-cashpay/merchant/services/category"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	Pagination(c echo.Context) error
	GetById(c echo.Context) error
	BulkDelete(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
}

type domainHandler struct {
	serviceCategory s.CategoryService
}

func NewCategoryHandler(service s.CategoryService) DomainHandler {
	return &domainHandler{
		serviceCategory: service,
	}
}

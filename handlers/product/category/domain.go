package category

import (
	"github.com/labstack/echo/v4"
	s "github.com/srv-cashpay/merchant/services/product/category"
)

type DomainHandler interface {
	Get(c echo.Context) error
}

type domainHandler struct {
	serviceGetCategory s.GetCategoryService
}

func NewCategoryHandler(service s.GetCategoryService) DomainHandler {
	return &domainHandler{
		serviceGetCategory: service,
	}
}

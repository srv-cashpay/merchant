package subscribe

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
)

func (h *domainHandler) CreatePaypalOrder(c echo.Context) error {
	var req dto.PaypalCreateRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	res, err := h.serviceSubscribe.CreatePaypalOrder(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *domainHandler) CapturePaypalOrder(c echo.Context) error {
	orderID := c.Param("order_id")
	if orderID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Order ID is required")
	}

	res, err := h.serviceSubscribe.CapturePaypalOrder(orderID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

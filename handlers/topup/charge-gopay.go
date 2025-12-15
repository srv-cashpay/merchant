package topup

import (
	"net/http"

	dto "github.com/srv-cashpay/merchant/dto"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) ChargeGopay(c echo.Context) error {
	var req dto.ChargeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
	}

	response, err := h.serviceSubscribe.ChargeGopay(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

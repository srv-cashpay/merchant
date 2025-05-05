package subscribe

import (
	"github.com/labstack/echo/v4"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *domainHandler) CheckTransactionStatus(c echo.Context) error {
	orderID := c.Param("order_id")
	result, err := h.serviceSubscribe.CheckTransactionStatus(orderID)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

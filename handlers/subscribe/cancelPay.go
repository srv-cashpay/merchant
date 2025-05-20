package subscribe

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *domainHandler) CancelPay(c echo.Context) error {
	var req dto.GetorderID

	orderid, err := res.IsNumber(c, "order_id")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	req.OrderID = orderid
	response, status, err := h.serviceSubscribe.CancelPay(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(status, map[string]interface{}{"response": string(response)})
}

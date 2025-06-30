package order

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *domainHandler) Order(c echo.Context) error {
	var req dto.OrderRequest
	var resp dto.OrderResponse

	merchantID := c.QueryParam("merchant_id")
	if merchantID == "" {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	req.MerchantID = merchantID

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	resp, err = h.serviceOrder.Order(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)

}

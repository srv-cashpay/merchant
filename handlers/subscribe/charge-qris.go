package subscribe

import (
	"net/http"

	dto "github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) ChargeQris(c echo.Context) error {
	var req dto.ChargeRequest
	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	merchantId, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.MerchantID = merchantId
	req.UserID = userid
	req.CreatedBy = createdBy

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
	}

	resp, err := h.serviceSubscribe.ChargeQris(req)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

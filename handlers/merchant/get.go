package merchant

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) Get(c echo.Context) error {
	var req dto.GetMerchantRequest
	var resp dto.GetMerchantResponse

	userId, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = userId

	resp, err := b.serviceMerchant.Get(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}

package history

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/merchant/helpers"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) Get(c echo.Context) error {
	paginationDTO := helpers.GeneratePaginationRequest(c)

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	merchantId, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	paginationDTO.MerchantID = merchantId
	paginationDTO.UserID = userid

	if err := c.Bind(&paginationDTO); err != nil {
		return c.JSON(400, "Invalid request")
	}

	users := b.serviceHistory.Get(c, paginationDTO)

	return c.JSON(200, users)
}

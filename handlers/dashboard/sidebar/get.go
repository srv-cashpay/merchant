package sidebar

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) Get(c echo.Context) error {
	var req dto.SidebarRequest

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	merchantId, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.MerchantID = merchantId
	req.UserID = userid

	products, err := b.serviceSidebar.Get(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return c.JSON(http.StatusOK, products)
}
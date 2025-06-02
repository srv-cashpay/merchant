package role

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/helpers"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) Get(c echo.Context) error {
	var req dto.RoleUser
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

	products, err := b.serviceRole.Get(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return c.JSON(http.StatusOK, products)
}

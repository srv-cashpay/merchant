package deleteaccount

import (
	"strconv"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) Delete(c echo.Context) error {
	var req dto.DeleteDeleteAccountRequest
	deletedBy, ok := c.Get("DeletedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	req.DeletedBy = deletedBy
	idUint, err := res.IsNumber(c, "id")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	idInt, err := strconv.Atoi(idUint)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	req.ID = uint(idInt)

	data, err := b.serviceDeleteAccount.Delete(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

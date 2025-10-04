package voucher

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) GetVerifikasi(c echo.Context) error {
	var req dto.GetVerifikasi

	// idStr := c.Param("id")

	idUint, err := res.IsNumber(c, "id")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	req.ID = idUint

	transaction, err := b.serviceVoucher.GetVerifikasi(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)

	}

	return res.SuccessResponse(transaction).Send(c)

}

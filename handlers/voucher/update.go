package voucher

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) Update(c echo.Context) error {
	var req dto.VoucherUpdateRequest
	var resp dto.VoucherUpdateResponse

	req.ID = c.Param("id")
	req.MerchantID = c.Param("merchant_id")

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	resp, err = b.serviceVoucher.Update(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)

}

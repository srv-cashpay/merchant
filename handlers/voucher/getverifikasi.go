package voucher

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) GetVerifikasi(c echo.Context) error {
	var req dto.GetVerifikasi

	req.ID = c.Param("id")
	req.MerchantID = c.Param("merchant_id")

	transaction, err := b.serviceVoucher.GetVerifikasi(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)
	}

	return res.SuccessResponse(transaction).Send(c)
}

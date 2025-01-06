package merk

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) Get(c echo.Context) error {
	var req dto.MerkRequest
	var resp []dto.MerkResponse

	merchantId, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.MerchantID = merchantId
	req.CreatedBy = createdBy

	resp, err := b.serviceGetMerk.Get(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

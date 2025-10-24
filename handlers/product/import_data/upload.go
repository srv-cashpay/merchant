package import_data

import (
	"net/http"

	res "github.com/srv-cashpay/util/s/response"

	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/merchant/dto"
)

func (h *domainHandler) UploadProducts(c echo.Context) error {
	var req dto.ProductRequest

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "file wajib diunggah",
		})
	}

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
	req.UserID = userid
	req.MerchantID = merchantId
	req.CreatedBy = createdBy

	result, err := h.serviceImport.ImportProducts(c.Request().Context(), fileHeader)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

package import_data

import (
	"context"
	"net/http"

	res "github.com/srv-cashpay/util/s/response"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) UploadProducts(c echo.Context) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "file wajib diunggah",
		})
	}

	userID, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	merchantID, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	// 🔹 Kirim context + metadata ke service
	ctx := context.WithValue(c.Request().Context(), "UserId", userID)
	ctx = context.WithValue(ctx, "CreatedBy", createdBy)
	ctx = context.WithValue(ctx, "MerchantId", merchantID)

	result, err := h.serviceImport.ImportProducts(ctx, fileHeader)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

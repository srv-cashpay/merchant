package export_data

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *domainHandler) ExportExcel(c echo.Context) error {
	var req dto.ExportFilter

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	merchantId, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = userid
	req.MerchantID = merchantId

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}

	// 🔹 panggil service
	f, err := h.serviceExport.ExportExcel(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// 🔹 simpan ke buffer dulu
	buf, err := f.WriteToBuffer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate excel"})
	}

	filename := fmt.Sprintf("export_%s.xlsx", time.Now().Format("20060102_150405"))

	c.Response().Header().Set(echo.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+filename)
	c.Response().WriteHeader(http.StatusOK)

	// ✅ kirim biner data
	_, err = c.Response().Write(buf.Bytes())
	return err
}

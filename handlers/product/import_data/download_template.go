package import_data

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) DownloadTemplate(c echo.Context) error {
	data, filename, err := h.serviceImport.GenerateTemplate()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+filename)
	c.Response().Header().Set(echo.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	return c.Blob(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data)
}

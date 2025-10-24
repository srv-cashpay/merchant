package export_data

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) ExportExcel(c echo.Context) error {
	f, err := h.serviceExport.ExportExcel(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Simpan ke buffer dan kirim sebagai response download
	filename := "users.xlsx"
	c.Response().Header().Set(echo.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+filename)

	return f.Write(c.Response())
}

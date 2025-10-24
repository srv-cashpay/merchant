package export_data

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ExportFilter struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (h *domainHandler) ExportExcel(c echo.Context) error {
	var filter ExportFilter
	if err := c.Bind(&filter); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}

	f, err := h.serviceExport.ExportExcel(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	filename := "users.xlsx"
	c.Response().Header().Set(echo.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+filename)

	return f.Write(c.Response())
}

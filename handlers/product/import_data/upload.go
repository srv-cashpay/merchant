package import_data

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) UploadProducts(c echo.Context) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "file wajib diunggah",
		})
	}

	result, err := h.serviceImport.ImportProducts(c.Request().Context(), fileHeader)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

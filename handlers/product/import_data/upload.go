package import_data

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) UploaProdducts(c echo.Context) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "file wajib diunggah"})
	}

	// misal user ID dari JWT
	userID := c.Get("user_id").(string)

	result, err := h.serviceImport.ImportProducts(c.Request().Context(), fileHeader, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

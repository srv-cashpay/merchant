package history

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) CheckExpire(c echo.Context) error {
	orderID := c.Param("order_id")

	result, err := h.serviceHistory.CheckAndExpire(orderID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"order_id": result.OrderID,
		"status":   result.Status,
	})
}

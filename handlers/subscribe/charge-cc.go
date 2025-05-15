package subscribe

import (
	"log"
	"net/http"

	dto "github.com/srv-cashpay/merchant/dto"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) TokenizeCardHandler(c echo.Context) error {
	var req dto.TokenizeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Proses tokenisasi kartu
	transaction, err := h.serviceSubscribe.TokenizeCard(req)
	if err != nil {
		log.Println("Error during card tokenization:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Tokenization failed"})
	}

	return c.JSON(http.StatusOK, transaction)
}

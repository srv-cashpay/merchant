package subscribe

import (
	"log"
	"net/http"

	dto "github.com/srv-cashpay/merchant/dto"
	res "github.com/srv-cashpay/util/s/response"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) TokenizeCardHandler(c echo.Context) error {
	var req dto.TokenizeRequest
	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = userid
	req.CreatedBy = createdBy

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

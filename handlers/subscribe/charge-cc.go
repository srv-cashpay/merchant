package subscribe

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	dto "github.com/srv-cashpay/merchant/dto"

	"github.com/labstack/echo/v4"
)

// func (h *domainHandler) TokenizeCardHandler(c echo.Context) error {
// 	var req dto.TokenizeRequest
// 	userid, ok := c.Get("UserId").(string)
// 	if !ok {
// 		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
// 	}
// 	createdBy, ok := c.Get("CreatedBy").(string)
// 	if !ok {
// 		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
// 	}

// 	req.UserID = userid
// 	req.CreatedBy = createdBy

// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
// 	}

// 	// Proses tokenisasi kartu
// 	transaction, err := h.serviceSubscribe.TokenizeCard(req)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{
// 			"error": err.Error(),
// 		})
// 	}
// 	return c.JSON(http.StatusOK, transaction)
// }

// handler.go
func (h *domainHandler) TokenizeCardHandler(c echo.Context) error {
	var req dto.TokenizeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid payload"})
	}

	form := url.Values{}
	form.Set("card_number", req.CardNumber)
	form.Set("card_exp_month", req.ExpiryMonth)
	form.Set("card_exp_year", req.ExpiryYear)
	form.Set("card_cvv", req.CVV)
	form.Set("client_key", os.Getenv("MIDTRANS_CLIENT_KEY"))

	httpReq, _ := http.NewRequest("POST", "https://api.sandbox.midtrans.com/v2/token", strings.NewReader(form.Encode()))
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to contact Midtrans"})
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if resp.StatusCode != 200 {
		return c.JSON(resp.StatusCode, result)
	}

	// Optional: langsung charge ke Midtrans
	// h.ChargeWithToken(result["token_id"].(string), req.OrderID, req.Amount)

	return c.JSON(http.StatusOK, result)
}

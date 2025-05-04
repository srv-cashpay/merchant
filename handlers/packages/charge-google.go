package packages

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
)

func (h *domainHandler) ChargeGpay(c echo.Context) error {
	var req dto.ChargeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
	}

	if req.OrderID == "" || req.Amount <= 0 || req.CardToken == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Missing required fields: order_id, amount, or card_token",
		})
	}

	// Mengirim request menggunakan Google Pay dengan token
	payload := map[string]interface{}{
		"payment_type": "google_pay",
		"transaction_details": map[string]interface{}{
			"order_id":     req.OrderID,
			"gross_amount": req.Amount,
		},
		"google_pay": map[string]interface{}{
			"token_id": req.CardToken, // Token Google Pay yang diterima dari frontend
		},
		"custom_expiry": map[string]interface{}{
			"order_time":      time.Now().Format("2006-01-02 15:04:05 -0700"),
			"expiry_duration": 1,
			"unit":            "hour",
		},
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to encode payload",
		})
	}

	httpReq, err := http.NewRequest("POST", dto.GetMidtransEndpoint(), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create HTTP request",
		})
	}

	auth := base64.StdEncoding.EncodeToString([]byte(dto.GetMidtransServerKey() + ":"))
	httpReq.Header.Set("Authorization", "Basic "+auth)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 20 * time.Second}
	res, err := client.Do(httpReq)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"error":   "Failed to contact Midtrans",
			"details": err.Error(),
		})
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to read Midtrans response",
		})
	}

	var parsed dto.GooglePayResponse
	if err := json.Unmarshal(resBody, &parsed); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   "Invalid response from Midtrans",
			"details": string(resBody),
		})
	}

	if parsed.StatusCode != "201" {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"error":   "Midtrans returned an error",
			"code":    parsed.StatusCode,
			"message": parsed.StatusMessage,
		})
	}

	return c.JSON(http.StatusOK, parsed)
}

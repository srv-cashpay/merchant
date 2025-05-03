package packages

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	MidtransServerKey = "SB-Mid-server-NYdY8JWbr0L0pc0f3jPXsLHH" // Ganti dengan server key kamu
	MidtransEndpoint  = "https://api.sandbox.midtrans.com/v2/charge"
)

type ChargeRequest struct {
	OrderID string `json:"order_id"`
	Amount  int64  `json:"amount"`
}

type VAResponse struct {
	OrderID           string `json:"order_id"`
	TransactionID     string `json:"transaction_id"`
	TransactionStatus string `json:"transaction_status"`
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	VANumbers         []struct {
		Bank     string `json:"bank"`
		VANumber string `json:"va_number"`
	} `json:"va_numbers"`
	ExpiryTime string `json:"expiry_time"`
}

func (h *domainHandler) Charge(c echo.Context) error {
	var req ChargeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
	}

	if req.OrderID == "" || req.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Missing required fields: order_id or amount",
		})
	}

	payload := map[string]interface{}{
		"payment_type": "bank_transfer",
		"transaction_details": map[string]interface{}{
			"order_id":     req.OrderID,
			"gross_amount": req.Amount,
		},
		"bank_transfer": map[string]interface{}{
			"bank": "bni",
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

	httpReq, err := http.NewRequest("POST", MidtransEndpoint, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create HTTP request",
		})
	}

	auth := base64.StdEncoding.EncodeToString([]byte(MidtransServerKey + ":"))
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

	var parsed VAResponse
	if err := json.Unmarshal(resBody, &parsed); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   "Invalid response from Midtrans",
			"details": string(resBody),
		})
	}

	// Cek status_code di dalam isi response
	if parsed.StatusCode != "201" {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"error":   "Midtrans returned an error",
			"code":    parsed.StatusCode,
			"message": parsed.StatusMessage,
		})
	}

	// Berhasil
	return c.JSON(http.StatusOK, parsed)
}

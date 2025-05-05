package subscribe

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *domainHandler) CheckTransactionStatus(c echo.Context) error {
	orderID := c.Param("order_id")

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY") // Simpan server key di env
	if serverKey == "" {
		return res.ErrorResponse(fmt.Errorf("Midtrans server key is not set")).Send(c)
	}

	url := fmt.Sprintf("https://api.sandbox.midtrans.com/v2/%s/status", orderID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	// Set headers
	auth := base64.StdEncoding.EncodeToString([]byte(serverKey + ":"))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")

	// Do request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

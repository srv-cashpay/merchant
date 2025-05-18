package subscribe

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	dto "github.com/srv-cashpay/merchant/dto"
)

func (r *subscribeRepository) ChargeGopay(req dto.ChargeRequest) (*dto.GopayResponse, error) {
	payload := map[string]interface{}{
		"payment_type": "gopay",
		"transaction_details": map[string]interface{}{
			"order_id":     req.OrderID,
			"gross_amount": req.GrossAmount,
		},
		"gopay": map[string]interface{}{
			"enable_callback": true,
			"callback_url":    "myapp://payment-callback?order_id=" + req.OrderID,
		},
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", dto.GetMidtransEndpoint(), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(dto.GetMidtransServerKey() + ":"))
	httpReq.Header.Set("Authorization", "Basic "+auth)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 20 * time.Second}
	res, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var parsed dto.GopayResponse
	if err := json.Unmarshal(resBody, &parsed); err != nil {
		return nil, errors.New("Invalid Midtrans response: " + string(resBody))
	}

	if parsed.StatusCode != "201" {
		return nil, errors.New("Midtrans error: " + parsed.StatusMessage)
	}

	return &parsed, nil
}

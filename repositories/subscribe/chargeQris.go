package subscribe

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/srv-cashpay/merchant/entity"
	util "github.com/srv-cashpay/util/s"

	dto "github.com/srv-cashpay/merchant/dto"
)

func (r *subscribeRepository) ChargeQris(req dto.ChargeRequest) (*dto.QrisResponse, error) {
	payload := map[string]interface{}{
		"payment_type": "qris",
		"transaction_details": map[string]interface{}{
			"order_id":     req.OrderID,
			"gross_amount": req.Amount,
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

	var parsed dto.QrisResponse
	if err := json.Unmarshal(resBody, &parsed); err != nil {
		return nil, err
	}

	// Simpan ke DB tanpa memeriksa statusCode
	var qrUrl string
	for _, action := range parsed.Actions {
		if action.Name == "generate-qr-code" {
			qrUrl = action.URL
			break
		}
	}

	tx := entity.Package{
		ID:              util.GenerateRandomString(),
		UserID:          req.UserID,
		CreatedBy:       req.CreatedBy,
		OrderID:         parsed.OrderID,
		TransactionID:   parsed.TransactionID,
		GrossAmount:     req.Amount,
		PaymentType:     parsed.PaymentType,
		Status:          parsed.TransactionStatus,
		TransactionTime: parseTime(parsed.TransactionTime),
		Url:             qrUrl,
	}

	if err := r.DB.Create(&tx).Error; err != nil {
		return nil, err
	}

	return &parsed, nil
}

func parseTime(str string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return time.Now()
	}
	return t
}

package subscribe

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	util "github.com/srv-cashpay/util/s"
)

func (r *subscribeRepository) TokenizeCard(req dto.TokenizeRequest) (*dto.TokenizeResponse, error) {
	// Siapkan data request
	requestBody := map[string]interface{}{
		"payment_type": "credit_card",
		"transaction_details": map[string]interface{}{
			"order_id":     req.OrderID,
			"gross_amount": req.Amount,
		},
		"credit_card": map[string]interface{}{
			"number":       req.CardNumber,
			"expiry_month": req.ExpiryMonth,
			"expiry_year":  req.ExpiryYear,
			"cvv":          req.CVV,
		},
	}

	// Encode ke JSON
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Buat request
	httpReq, err := http.NewRequest("POST", dto.GetMidtransTokenize(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set header
	auth := base64.StdEncoding.EncodeToString([]byte(dto.GetMidtransServerKey() + ":"))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Basic "+auth)

	// Kirim request
	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.New("failed to contact Midtrans")
	}
	defer resp.Body.Close()

	// Parse response
	var response dto.TokenizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	// Jika ada error dari API
	// if resp.StatusCode != 200 {
	// 	return nil, fmt.Errorf("failed to tokenize card: %s", response.Status)
	// }

	tx := entity.CreditCard{
		ID:            util.GenerateRandomString(),
		UserID:        req.UserID,
		CreatedBy:     req.CreatedBy,
		OrderID:       response.OrderID,
		TransactionID: response.TransactionID,
		CardNumber:    req.CardNumber,
		ExpiryMonth:   req.ExpiryMonth,
		ExpiryYear:    req.ExpiryYear,
		CVV:           req.CVV,
		Amount:        req.Amount,
	}

	if err := r.DB.Create(&tx).Error; err != nil {
		return nil, err
	}

	return &response, nil
}

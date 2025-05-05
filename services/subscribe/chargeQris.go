package subscribe

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *subscribeService) ChargeQris(req dto.ChargeRequest) (*dto.QrisResponse, error) {
	if req.OrderID == "" || req.Amount <= 0 {
		return nil, errors.New("missing required fields: order_id or amount")
	}

	resp, err := s.Repo.ChargeQris(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != "201" {
		return nil, errors.New(resp.StatusMessage)
	}

	return resp, nil
}

func (s *subscribeService) CheckTransactionStatus(orderID string) (map[string]interface{}, error) {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		return nil, fmt.Errorf("midtrans server key is not set")
	}

	url := fmt.Sprintf("https://api.sandbox.midtrans.com/v2/%s/status", orderID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(serverKey + ":"))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	transactionStatus := result["transaction_status"].(string)

	// Update DB status
	if err := s.Repo.UpdateStatus(orderID, transactionStatus); err != nil {
		return nil, err
	}

	// Jika berhasil dibayar, aktifkan akun
	if transactionStatus == "settlement" {
		if err := s.Repo.UpdateUserVerified(orderID); err != nil {
			return nil, err
		}
	}

	return result, nil
}

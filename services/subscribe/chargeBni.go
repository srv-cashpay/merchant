package subscribe

import (
	"errors"
	"time"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (s *subscribeService) ChargeBni(req dto.ChargeRequest) (*dto.VAResponse, error) {
	if req.OrderID == "" || req.GrossAmount <= 0 {
		return nil, errors.New("missing required fields: order_id or amount")
	}

	// CEK APAKAH ORDER_ID SUDAH ADA DI DATABASE
	var existing entity.Subscribe
	err := s.Repo.FindByOrderID(req.OrderID, req.UserID, &existing)
	if err == nil {
		if existing.Status == "pending" && time.Now().Before(existing.TransactionTime.Add(1*time.Hour)) {
			// Masih dalam masa berlaku
			return nil, errors.New("Transaksi dengan order_id ini masih aktif. Silakan selesaikan pembayaran sebelumnya.")
		} else {
			// Sudah expired atau status bukan pending
			return nil, errors.New("Order ID sudah digunakan sebelumnya. Silakan buat transaksi baru.")
		}
	}

	resp, err := s.Repo.ChargeBni(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != "201" {
		return nil, errors.New(resp.StatusMessage)
	}

	return resp, nil
}

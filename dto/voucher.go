package dto

import (
	"time"
)

type VoucherRequest struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	Nomor       string    `json:"nomor"`
	VoucherName string    `json:"voucher_name"`
	VoucherLink string    `json:"voucher_link"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	Status      int       `json:"status"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

package dto

import "os"

type CreateTransactionRequest struct {
	OrderID     string  `json:"order_id"`
	UserID      string  `json:"user_id"`
	CreatedBy   string  `json:"created_by"`
	GrossAmount float64 `json:"gross_amount"`
	PaymentType string  `json:"payment_type"` // misalnya: gopay, bca_va, etc.
}

type TransactionStatusResponse struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
	PaymentType       string `json:"payment_type"`
	GrossAmount       string `json:"gross_amount"`
	TransactionTime   string `json:"transaction_time"`
	TransactionID     string `json:"transaction_id"`
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	FraudStatus       string `json:"fraud_status,omitempty"`
	SettlementTime    string `json:"settlement_time,omitempty"`
}

type ChargeRequest struct {
	OrderID   string `json:"order_id"`
	Amount    int64  `json:"amount"`
	CardToken string `json:"card_token"` // Token kartu kredit yang didapatkan dari frontend
	Channel   string `json:"channel"`    // Metode pembayaran yang digunakan, contoh: "credit_card"
	UserID    string `json:"user_id"`
	CreatedBy string `json:"created_by"`
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

type QrisChargeRequest struct {
	PaymentType        string             `json:"payment_type"` // "qris"
	TransactionDetails TransactionDetails `json:"transaction_details"`
	ItemDetails        []ItemDetail       `json:"item_details,omitempty"`
	CustomerDetails    *CustomerDetails   `json:"customer_details,omitempty"`
	CustomExpiry       *CustomExpiry      `json:"custom_expiry,omitempty"`
}

type TransactionDetails struct {
	OrderID     string  `json:"order_id"`
	GrossAmount float64 `json:"gross_amount"`
}

type ItemDetail struct {
	ID       string  `json:"id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Name     string  `json:"name"`
}

type CustomerDetails struct {
	FirstName       string   `json:"first_name"`
	LastName        string   `json:"last_name"`
	Email           string   `json:"email"`
	Phone           string   `json:"phone"`
	BillingAddress  *Address `json:"billing_address,omitempty"`
	ShippingAddress *Address `json:"shipping_address,omitempty"`
}

type Address struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	CountryCode string `json:"country_code"`
}

type CustomExpiry struct {
	OrderTime      string `json:"order_time"`      // format: "2006-01-02 15:04:05 -0700"
	ExpiryDuration int    `json:"expiry_duration"` // in minutes or hours
	Unit           string `json:"unit"`            // "minute", "hour", or "day"
}

type QrisResponse struct {
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	TransactionID     string `json:"transaction_id"`
	OrderID           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	PaymentType       string `json:"payment_type"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	Actions           []struct {
		Name   string `json:"name"`
		Method string `json:"method"`
		URL    string `json:"url"`
	} `json:"actions"`
}

type GooglePayResponse struct {
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	TransactionID     string `json:"transaction_id"`
	OrderID           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	PaymentType       string `json:"payment_type"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	Actions           []struct {
		Name   string `json:"name"`
		Method string `json:"method"`
		URL    string `json:"url"`
	} `json:"actions"`
}

type CardResponse struct {
	StatusCode        string `json:"status_code"`        // Kode status dari transaksi (misal: "201" jika sukses)
	StatusMessage     string `json:"status_message"`     // Pesan status transaksi
	TransactionID     string `json:"transaction_id"`     // ID transaksi yang unik
	OrderID           string `json:"order_id"`           // ID pesanan
	GrossAmount       string `json:"gross_amount"`       // Jumlah transaksi (gross amount)
	PaymentType       string `json:"payment_type"`       // Jenis pembayaran (misalnya: "credit_card")
	TransactionTime   string `json:"transaction_time"`   // Waktu transaksi dilakukan
	TransactionStatus string `json:"transaction_status"` // Status transaksi (capture, pending, deny)
	FraudStatus       string `json:"fraud_status"`       // Status verifikasi penipuan (accept, challenge, deny)
	RedirectURL       string `json:"redirect_url"`       // URL untuk otentikasi lebih lanjut (jika diperlukan 3D Secure)
}

type GopayResponse struct {
	StatusCode        string        `json:"status_code"`
	StatusMessage     string        `json:"status_message"`
	TransactionID     string        `json:"transaction_id"`
	OrderID           string        `json:"order_id"`
	GrossAmount       string        `json:"gross_amount"`
	PaymentType       string        `json:"payment_type"`
	TransactionTime   string        `json:"transaction_time"`
	TransactionStatus string        `json:"transaction_status"`
	Actions           []GopayAction `json:"actions,omitempty"`
	FraudStatus       string        `json:"fraud_status,omitempty"`
	GopayPaymentURL   string        `json:"gopay_payment_url,omitempty"`
}

type ShopeePayResponse struct {
	StatusCode        string        `json:"status_code"`
	StatusMessage     string        `json:"status_message"`
	TransactionID     string        `json:"transaction_id"`
	OrderID           string        `json:"order_id"`
	GrossAmount       string        `json:"gross_amount"`
	PaymentType       string        `json:"payment_type"`
	TransactionTime   string        `json:"transaction_time"`
	TransactionStatus string        `json:"transaction_status"`
	FraudStatus       string        `json:"fraud_status,omitempty"`
	Actions           []GopayAction `json:"actions,omitempty"` // Sama dengan GoPay, bisa reuse struct
}

// Untuk menampung informasi actions (seperti deeplink, QR, dsb)
type GopayAction struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	URL    string `json:"url"`
}

func GetMidtransEndpoint() string {
	val := os.Getenv("MIDTRANS_ENDPOINT")
	if val == "" {
		return "https://api.sandbox.midtrans.com/v2/charge"
	}
	return val
}

func GetMidtransServerKey() string {
	val := os.Getenv("MIDTRANS_SERVER_KEY")
	if val == "" {
		return "your-default-midtrans-server-key"
	}
	return val
}

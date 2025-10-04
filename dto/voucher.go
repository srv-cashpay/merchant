package dto

type VoucherRequest struct {
	ID              string            `json:"id"`
	UserID          string            `json:"user_id"`
	MerchantID      string            `json:"merchant_id"`
	Nomor           int               `json:"nomor"`
	VoucherName     string            `json:"voucher_name"`
	StartDate       string            `json:"start_date"`
	EndDate         string            `json:"end_date"`
	Status          bool              `json:"status"`
	CreatedBy       string            `json:"created_by"`
	VoucherGenerate []VoucherGenerate `json:"voucher_generate"`
}

type VoucherResponse struct {
	ID              string `json:"id"`
	UserID          string `json:"user_id"`
	MerchantID      string `json:"merchant_id"`
	VoucherGenerate []VoucherGenerate
	CreatedBy       string `json:"created_by"`
}

type VoucherGenerate struct {
	ID          string `json:"id"`
	MerchantID  string `json:"merchant_id"`
	VoucherName string `json:"voucher_name"`
	VoucherLink string `json:"voucher_link"`
	VoucherQR   string `json:"voucher_qr"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Status      bool   `json:"status"`
}

type VoucherGetByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type GetVerifikasi struct {
	ID         string `param:"id" validate:"required"`
	MerchantID string `param:"merchant_id" validate:"required"`
}

type GetVerifikasiResponse struct {
	VoucherGenerate []VoucherGenerate `json:"voucher_generate"`
}

type VoucherDeleteRequest struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type VoucherDeleteResponse struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type VoucherUpdateRequest struct {
	ID         string `param:"id" validate:"required"`
	UserID     string `json:"user_id"`
	MerchantID string `param:"merchant_id" validate:"required"`
	Status     bool   `json:"status"`
	UpdatedBy  string `json:"updated_by"`
}

type VoucherUpdateResponse struct {
	Status bool `json:"status"`
}

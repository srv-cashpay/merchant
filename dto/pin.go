package dto

type PinRequest struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	MerchantID  string `json:"merchant_id"`
	Pin         int    `json:"pin"`
	CreatedBy   string `json:"created_by"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type PinResponse struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	MerchantID  string `json:"merchant_id"`
	Pin         int    `json:"pin"`
	CreatedBy   string `json:"created_by"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type GetByIdPinRequest struct {
	ID string `param:"id" validate:"required"`
}

type DeletePinRequest struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeletePinResponse struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type PinUpdateRequest struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	MerchantID  string `json:"merchant_id"`
	Pin         int    `json:"pin"`
	UpdatedBy   string `json:"updated_by"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type PinUpdateResponse struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	MerchantID  string `json:"merchant_id"`
	Pin         int    `json:"pin"`
	UpdatedBy   string `json:"updated_by"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type BulkDeletePinRequest struct {
	ID        []string `json:"id"`
	DeletedBy string   `json:"deleted_by"`
}

type BulkDeletePinResponse struct {
	ID        []string `json:"id"`
	DeletedBy string   `json:"deleted_by"`
	Count     int      `json:"count"`
}

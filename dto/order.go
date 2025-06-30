package dto

type OrderRequest struct {
	ID          uint   `json:"id"`
	UserID      string `json:"user_id"`
	MerchantID  string `json:"merchant_id"`
	OrderName   string `json:"order_name"`
	CreatedBy   string `json:"created_by"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type OrderResponse struct {
	ID          uint   `json:"id"`
	UserID      string `json:"user_id"`
	MerchantID  string `json:"merchant_id"`
	OrderName   string `json:"order_name"`
	CreatedBy   string `json:"created_by"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type OrderUpdateRequest struct {
	ID          uint   `json:"id"`
	UserID      string `json:"user_id"`
	MerchantID  string `json:"merchant_id"`
	OrderName   string `json:"order_name"`
	UpdatedBy   string `json:"updated_by"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type OrderUpdateResponse struct {
	ID          uint   `json:"id"`
	UserID      string `json:"user_id"`
	MerchantID  string `json:"merchant_id"`
	OrderName   string `json:"order_name"`
	UpdatedBy   string `json:"updated_by"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type GetByIdOrderRequest struct {
	ID uint `param:"id" validate:"required"`
}

type DeleteOrderRequest struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeleteOrderResponse struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

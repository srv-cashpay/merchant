package dto

type OrderRequest struct {
	ID          uint          `json:"id"`
	UserID      string        `json:"user_id"`
	MerchantID  string        `json:"merchant_id"`
	Product     []ProductItem `json:"product"`
	ProductJSON string        `json:"-" gorm:"column:product"`
	OrderName   string        `json:"order_name"`
	CreatedBy   string        `json:"created_by"`
	Description string        `json:"description"`
	Status      int           `json:"status"`
}

type ProductItem struct {
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type OrderResponse struct {
	ID          uint          `json:"id"`
	UserID      string        `json:"user_id"`
	MerchantID  string        `json:"merchant_id"`
	OrderName   string        `json:"order_name"`
	CreatedBy   string        `json:"created_by"`
	Description string        `json:"description"`
	Status      int           `json:"status"`
	Product     []ProductItem `json:"product"`
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

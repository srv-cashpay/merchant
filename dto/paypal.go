package dto

type CapturePaypalOrderRequest struct {
	OrderID string `json:"order_id" validate:"required"`
}

type PaypalCreateRequest struct {
	Amount   string `json:"amount" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}

type PaypalOrderResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Link   string `json:"link"`
}

type PaypalCaptureResponse struct {
	Status string `json:"status"`
}

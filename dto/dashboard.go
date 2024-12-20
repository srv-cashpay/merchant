package dto

type GetDashboardRequest struct {
	ID                    string `param:"id" validate:"required"`
	MerchantID            string `json:"merchant_id"`
	CreatedBy             string `json:"created_by"`
	TotalProductsActive   int64  `json:"total_products_active"`
	TotalProductsInactive int64  `json:"total_products_inactive"`
	TotalPrice            int64  `json:"total_price"`
	TotalWaiting          int64  `json:"total_waiting"`
	PaymentCancel         int64  `json:"payment_cancel"`
}

type GetDashboardResponse struct {
	TotalProductsActive   int64               `json:"total_products_active"`
	TotalProductsInactive int64               `json:"total_products_inactive"`
	TotalSales            int64               `json:"total_sales"`
	TotalPrice            int64               `json:"total_price"`
	WaitingPayment        int64               `json:"total_waiting_payment"`
	TotalWaiting          int64               `json:"total_waiting"`
	PaymentCancel         int64               `json:"payment_cancel"`
	ProductPercentages    []ProductPercentage `json:"product_percentages"` // Persentase per produk
}

type ProductPercentage struct {
	ProductID         string  `json:"product_id"`
	ProductName       string  `json:"product_name"`
	TotalLunas        int64   `json:"total_lunas"`
	TotalTransactions int64   `json:"total_transactions"`
	Percentage        float64 `json:"percentage"`
}

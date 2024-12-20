package dashboard

import (
	"github.com/srv-cashpay/merchant/dto"
	entityPos "github.com/srv-cashpay/pos/entity"
	"github.com/srv-cashpay/product/entity"
)

func (r *dashboardRepository) Get(req dto.GetDashboardRequest) (dto.GetDashboardResponse, error) {
	var response dto.GetDashboardResponse

	// Hitung total produk aktif
	if err := r.DB.Model(&entity.Product{}).
		Where("merchant_id = ? AND status = ?", req.MerchantID, 1).
		Count(&req.TotalProductsActive).Error; err != nil {
		return dto.GetDashboardResponse{}, err
	}
	response.TotalProductsActive = req.TotalProductsActive

	// Hitung total produk tidak aktif
	if err := r.DB.Model(&entity.Product{}).
		Where("merchant_id = ? AND status = ?", req.MerchantID, 2).
		Count(&req.TotalProductsInactive).Error; err != nil {
		return dto.GetDashboardResponse{}, err
	}
	response.TotalProductsInactive = req.TotalProductsInactive

	var totalSales int64
	if err := r.DB.Model(&entityPos.Pos{}).
		Where("merchant_id = ?", req.MerchantID).
		Count(&totalSales).Error; err != nil {
		return dto.GetDashboardResponse{}, err
	}
	response.TotalSales = totalSales

	if err := r.DB.Raw(`
		SELECT COALESCE(SUM((json_data->>'price')::int * (json_data->>'quantity')::int), 0) 
		FROM pos,
		LATERAL json_array_elements(pos.product) AS json_data
		WHERE pos.merchant_id = ? AND pos.deleted_at IS NULL 
		AND pos.status_payment IN ('Lunas')
	`, req.MerchantID).Scan(&req.TotalPrice).Error; err != nil {
		return response, err
	}
	response.TotalPrice = req.TotalPrice

	if err := r.DB.Raw(`
		SELECT COALESCE(SUM((json_data->>'price')::int * (json_data->>'quantity')::int), 0) 
		FROM pos,
		LATERAL json_array_elements(pos.product) AS json_data
		WHERE pos.merchant_id = ? AND pos.deleted_at IS NULL 
		AND pos.status_payment IN ('Menunggu Pembayaran')
	`, req.MerchantID).Scan(&req.TotalWaiting).Error; err != nil {
		return response, err
	}
	response.TotalWaiting = req.TotalWaiting

	var waitingPayment int64
	if err := r.DB.Model(&entityPos.Pos{}).
		Where("merchant_id = ? AND status_payment = ?", req.MerchantID, "Menunggu Pembayaran").
		Count(&waitingPayment).Error; err != nil {
		return dto.GetDashboardResponse{}, err
	}
	response.WaitingPayment = waitingPayment

	if err := r.DB.Model(&entityPos.Pos{}).
		Where("merchant_id = ? AND status_payment = ?", req.MerchantID, "Batal Pembayaran").
		Count(&req.PaymentCancel).Error; err != nil {
		return dto.GetDashboardResponse{}, err
	}
	response.PaymentCancel = req.PaymentCancel

	var productPercentages []dto.ProductPercentage
	if err := r.DB.Raw(`
    SELECT 
        json_data->>'product_id' AS product_id,
        json_data->>'product_name' AS product_name,
        COUNT(*) AS total_transactions,
        SUM(CASE WHEN pos.status_payment = 'Lunas' THEN 1 ELSE 0 END) AS total_lunas,
        (SUM(CASE WHEN pos.status_payment = 'Lunas' THEN 1 ELSE 0 END) * 100.0 / COUNT(*)) AS percentage
    FROM pos,
    LATERAL json_array_elements(pos.product) AS json_data
    WHERE pos.merchant_id = ? 
    AND pos.deleted_at IS NULL
    GROUP BY json_data->>'product_id', json_data->>'product_name'
`, req.MerchantID).Scan(&productPercentages).Error; err != nil {
		return response, err
	}

	response.ProductPercentages = productPercentages

	return response, nil
}

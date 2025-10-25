package export_data

import (
	"context"

	dto "github.com/srv-cashpay/merchant/dto"

	"github.com/srv-cashpay/product/entity"
)

func (r *exportRepository) FindAllByFilter(ctx context.Context, req dto.ExportFilter) ([]entity.Product, error) {
	var products []entity.Product

	tx := r.DB.WithContext(ctx).Where("merchant_id = ?", req.MerchantID)

	if req.From != "" && req.To != "" {
		tx = tx.Where("created_at BETWEEN ? AND ?", req.From, req.To)
	}

	if err := tx.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

package export_data

import (
	"context"

	"github.com/srv-cashpay/product/entity"
)

func (r *exportRepository) FindAll(ctx context.Context) ([]entity.Product, error) {
	var products []entity.Product
	if err := r.DB.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

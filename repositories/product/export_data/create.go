package export_data

import (
	"context"

	"github.com/srv-cashpay/product/entity"
)

func (r *exportRepository) FindAll(ctx context.Context) ([]entity.Product, error) {
	var users []entity.Product
	if err := r.DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

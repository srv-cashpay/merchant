package import_data

import (
	"context"

	"github.com/srv-cashpay/product/entity"
)

func (r *importRepository) BulkInsert(ctx context.Context, products []entity.Product) error {
	if len(products) == 0 {
		return nil
	}
	return r.DB.WithContext(ctx).Create(&products).Error
}

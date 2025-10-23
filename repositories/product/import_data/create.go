package import_data

import (
	"context"

	"github.com/srv-cashpay/product/entity"
)

func (r *importRepository) Create(ctx context.Context, product *entity.Product) error {
	return r.DB.WithContext(ctx).Create(product).Error
}

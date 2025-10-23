package import_data

import (
	"context"
	"errors"

	"github.com/srv-cashpay/product/entity"
)

func (r *importRepository) SaveBatch(ctx context.Context, products []entity.Product) error {
	if len(products) == 0 {
		return errors.New("tidak ada produk untuk disimpan")
	}

	tx := r.DB.WithContext(ctx).Create(&products)
	return tx.Error
}

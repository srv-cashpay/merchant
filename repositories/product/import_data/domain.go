package import_data

import (
	"context"

	"github.com/srv-cashpay/product/entity"
	"gorm.io/gorm"
)

type DomainRepository interface {
	BulkInsert(ctx context.Context, products []entity.Product) error
}

type importRepository struct {
	DB *gorm.DB
}

func NewImportRepository(DB *gorm.DB) DomainRepository {
	return &importRepository{
		DB: DB,
	}
}

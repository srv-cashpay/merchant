package import_data

import (
	"context"

	"github.com/srv-cashpay/product/entity"
	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(ctx context.Context, product *entity.Product) error
}

type importRepository struct {
	DB *gorm.DB
}

func NewImportRepository(DB *gorm.DB) DomainRepository {
	return &importRepository{
		DB: DB,
	}
}

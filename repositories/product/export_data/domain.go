package export_data

import (
	"context"

	"github.com/srv-cashpay/product/entity"
	"gorm.io/gorm"
)

type DomainRepository interface {
	FindAll(ctx context.Context) ([]entity.Product, error)
}

type exportRepository struct {
	DB *gorm.DB
}

func NewExportRepository(DB *gorm.DB) DomainRepository {
	return &exportRepository{
		DB: DB,
	}
}

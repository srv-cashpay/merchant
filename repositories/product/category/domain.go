package category

import (
	"github.com/srv-cashpay/merchant/entity"
	"gorm.io/gorm"
)

type DomainRepository interface {
	Get() ([]entity.Category, error)
}

type getcategoryRepository struct {
	DB *gorm.DB
}

func NewGetCategoryRepository(DB *gorm.DB) DomainRepository {
	return &getcategoryRepository{DB: DB}
}

package history

import (
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req *dto.Pagination) (RepositoryResult, int)
}

type historyRepository struct {
	DB *gorm.DB
}

func NewHistoryRepository(DB *gorm.DB) DomainRepository {
	return &historyRepository{
		DB: DB,
	}
}

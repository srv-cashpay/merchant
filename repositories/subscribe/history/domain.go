package history

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req *dto.Pagination) (RepositoryResult, int)
	CheckAndExpireIfNeeded(orderID string) (*entity.Subscribe, error)
	GetById(req dto.GetHistory) (*dto.VAResponse, error)
}

type historyRepository struct {
	DB *gorm.DB
}

func NewHistoryRepository(DB *gorm.DB) DomainRepository {
	return &historyRepository{
		DB: DB,
	}
}

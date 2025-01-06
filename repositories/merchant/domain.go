package merchant

import (
	"sync"

	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req dto.GetMerchantRequest) (dto.GetMerchantResponse, error)
	GetById(req dto.GetByIdRequest) (*dto.GetMerchantResponse, error)
	Update(req dto.UpdateMerchantRequest) (dto.UpdateMerchantResponse, error)
}

type merchantRepository struct {
	DB *gorm.DB
	mu sync.Mutex
}

func NewMerchantRepository(DB *gorm.DB) DomainRepository {
	return &merchantRepository{
		DB: DB,
	}
}

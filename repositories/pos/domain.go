package pos

import (
	"sync"

	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.PosRequest) (dto.PosResponse, error)
}

type posRepository struct {
	DB *gorm.DB
	mu sync.Mutex
}

func NewPosRepository(DB *gorm.DB) DomainRepository {
	return &posRepository{
		DB: DB,
	}
}

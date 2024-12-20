package authenticator_request

import (
	"sync"

	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.AuthenticatorRequest) (dto.AuthenticatorResponse, error)
}

type authenticatorRepository struct {
	DB *gorm.DB
	mu sync.Mutex
}

func NewAuthenticatorRepository(DB *gorm.DB) DomainRepository {
	return &authenticatorRepository{
		DB: DB,
	}
}

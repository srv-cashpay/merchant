package subscribe

import (
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.PackagesRequest) (dto.PackagesResponse, error)
	UpdateStatus(orderID string, status string) error
	UpdateUserVerified(orderID string) error
	ChargeQris(req dto.ChargeRequest) (*dto.QrisResponse, error)
	CheckStatus(orderID string) (map[string]interface{}, error)
	ChargeBca(req dto.ChargeRequest) (*dto.VAResponse, error)
	ChargeBri(req dto.ChargeRequest) (*dto.VAResponse, error)
	ChargeBni(req dto.ChargeRequest) (*dto.VAResponse, error)
}

type subscribeRepository struct {
	DB *gorm.DB
}

func NewSubscribeRepository(DB *gorm.DB) DomainRepository {
	return &subscribeRepository{
		DB: DB,
	}
}

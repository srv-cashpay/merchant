package qris

import (
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.CoQrisRequest) (dto.CoQrisResponse, error)
}

type qrisRepository struct {
	DB *gorm.DB
}

func NewQrisRepository(DB *gorm.DB) DomainRepository {
	return &qrisRepository{
		DB: DB,
	}
}

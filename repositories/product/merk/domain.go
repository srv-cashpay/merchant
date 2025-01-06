package merk

import (
	"github.com/srv-cashpay/merchant/dto"
	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req dto.MerkRequest) ([]dto.MerkResponse, error)
}

type getmerkRepository struct {
	DB *gorm.DB
}

func NewGetMerkRepository(DB *gorm.DB) DomainRepository {
	return &getmerkRepository{DB: DB}
}

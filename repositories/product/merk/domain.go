package merk

import (
	"github.com/srv-cashpay/merchant/entity"
	"gorm.io/gorm"
)

type DomainRepository interface {
	Get() ([]entity.Merk, error)
}

type getmerkRepository struct {
	DB *gorm.DB
}

func NewGetMerkRepository(DB *gorm.DB) DomainRepository {
	return &getmerkRepository{DB: DB}
}

package packages

import (
	auth "github.com/srv-cashpay/auth/entity"
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.PackagesRequest) (dto.PackagesResponse, error)
	UpdateStatus(orderID string, status string) error
	FindByID(id string) (auth.AccessDoor, error)
}

type packagesRepository struct {
	DB *gorm.DB
}

func NewPackagesRepository(DB *gorm.DB) DomainRepository {
	return &packagesRepository{
		DB: DB,
	}
}

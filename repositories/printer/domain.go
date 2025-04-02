package printer

import (
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.PrinterRequest) (dto.PrinterResponse, error)
	Get(req dto.GetPrinterRequest) (dto.GetPrinterResponse, error)
	GetById(req dto.GetByIdRequest) (*dto.GetPrinterResponse, error)
	Update(req dto.UpdatePrinterRequest) (dto.UpdatePrinterResponse, error)
}

type printerRepository struct {
	DB *gorm.DB
}

func NewPrinterRepository(DB *gorm.DB) DomainRepository {
	return &printerRepository{
		DB: DB,
	}
}

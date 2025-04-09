package printer

import (
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/printer"
)

type PrinterService interface {
	Create(req dto.PrinterRequest) (dto.PrinterResponse, error)
	Get(req dto.GetPrinterRequest) (dto.GetPrinterResponse, error)
	Update(req dto.UpdatePrinterRequest) (dto.UpdatePrinterResponse, error)
	Delete(req dto.DeletePrinterRequest) (dto.DeletePrinterResponse, error)
}

type printerService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPrinterService(Repo r.DomainRepository, jwtS m.JWTService) PrinterService {
	return &printerService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

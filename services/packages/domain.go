package packages

import (
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/packages"
)

type PackagesService interface {
	Create(req dto.PackagesRequest) (dto.PackagesResponse, error)
	UpdateStatus(orderID string, transactionStatus string) error
}

type packagesService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPackagesService(Repo r.DomainRepository, jwtS m.JWTService) PackagesService {
	return &packagesService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

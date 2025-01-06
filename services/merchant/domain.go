package product

import (
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/merchant"
)

type MerchantService interface {
	Get(req dto.GetMerchantRequest) (dto.GetMerchantResponse, error)
	Update(req dto.UpdateMerchantRequest) (dto.UpdateMerchantResponse, error)
}

type merchantService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewMerchantService(Repo r.DomainRepository, jwtS m.JWTService) MerchantService {
	return &merchantService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

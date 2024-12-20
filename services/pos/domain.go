package pos

import (
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/pos"
)

type PosService interface {
	Create(req dto.PosRequest) (dto.PosResponse, error)
}

type posService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPosService(Repo r.DomainRepository, jwtS m.JWTService) PosService {
	return &posService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

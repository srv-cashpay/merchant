package qris

import (
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/transactionmethode/qris"
)

type QrisService interface {
	Create(req dto.CoQrisRequest) (dto.CoQrisResponse, error)
}

type qrisService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewQrisService(Repo r.DomainRepository, jwtS m.JWTService) QrisService {
	return &qrisService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

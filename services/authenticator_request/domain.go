package authenticator_request

import (
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/authenticator_request"
)

type AuthenticatorService interface {
	Create(req dto.AuthenticatorRequest) (dto.AuthenticatorResponse, error)
}

type authenticatorService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewAuthenticatorService(Repo r.DomainRepository, jwtS m.JWTService) AuthenticatorService {
	return &authenticatorService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

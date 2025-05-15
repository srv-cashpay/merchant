package subscribe

import (
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/subscribe"
)

type SubscribeService interface {
	Create(req dto.SubscribeRequest) (dto.SubscribeResponse, error)
	UpdateStatus(orderID string, transactionStatus string) error
	ChargeQris(req dto.ChargeRequest) (*dto.QrisResponse, error)
	CheckTransactionStatus(request dto.CreateTransactionRequest) (dto.TransactionStatusResponse, error)
	ChargePermata(req dto.ChargeRequest) (*dto.VAPermataResponse, error)
	ChargeBri(req dto.ChargeRequest) (*dto.VAResponse, error)
	ChargeCimb(req dto.ChargeRequest) (*dto.VAResponse, error)
	ChargeBni(req dto.ChargeRequest) (*dto.VAResponse, error)
	CardPayment(cardData dto.TokenizeRequest) (*dto.TokenizeResponse, error)
}

type subscribeService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewSubscribeService(Repo r.DomainRepository, jwtS m.JWTService) SubscribeService {
	return &subscribeService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

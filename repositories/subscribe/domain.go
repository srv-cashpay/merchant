package subscribe

import (
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.SubscribeRequest) (dto.SubscribeResponse, error)
	UpdateStatus(orderID string, status string) error
	UpdateUserVerified(orderID string) error
	ChargeQris(req dto.ChargeRequest) (*dto.QrisResponse, error)
	CheckStatus(orderID string) (map[string]interface{}, error)
	ChargePermata(req dto.ChargeRequest) (*dto.VAPermataResponse, error)
	ChargeBri(req dto.ChargeRequest) (*dto.VAResponse, error)
	ChargeCimb(req dto.ChargeRequest) (*dto.VAResponse, error)
	ChargeBni(req dto.ChargeRequest) (*dto.VAResponse, error)
	CardPayment(cardData dto.CreditCardChargeRequest) (*dto.TokenizeResponse, error)
	ChargeGopay(req dto.ChargeRequest) (*dto.GopayResponse, error)
	CancelPay(req dto.GetorderID) ([]byte, int, error)
	UpdateSubscribeByOrderID(data dto.MidtransCancelResponse) error
}

type subscribeRepository struct {
	DB *gorm.DB
}

func NewSubscribeRepository(DB *gorm.DB) DomainRepository {
	return &subscribeRepository{
		DB: DB,
	}
}

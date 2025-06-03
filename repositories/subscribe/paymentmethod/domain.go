package paymentmethod

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.PaymentMethodRequest) (dto.PaymentMethodResponse, error)
	Get(req dto.PaymentMethodRequest) (dto.PaymentMethodResponse, error)
	GetById(req dto.GetByIdRequest) (*dto.PaymentMethodResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.PaymentMethodUpdateRequest) (dto.PaymentMethodUpdateResponse, error)
	CheckMerchantDetail(merchantID string, merchantDetail *entity.MerchantDetail) error
}

type paymentmethodRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) DomainRepository {
	return &paymentmethodRepository{
		DB: DB,
	}
}

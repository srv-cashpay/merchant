package payment

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.PaymentRequest) (dto.PaymentResponse, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetByIdRequest) (*dto.PaymentResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.PaymentUpdateRequest) (dto.PaymentUpdateResponse, error)
	CheckMerchantDetail(merchantID string, merchantDetail *entity.MerchantDetail) error
}

type paymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) DomainRepository {
	return &paymentRepository{
		DB: DB,
	}
}

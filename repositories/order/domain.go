package order

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.OrderRequest) (dto.OrderResponse, error)
	Order(req dto.OrderRequest) (dto.OrderResponse, error)
	Delete(req dto.DeleteOrderRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetByIdOrderRequest) (*dto.OrderResponse, error)
	Update(req dto.OrderUpdateRequest) (dto.OrderUpdateResponse, error)
	CheckMerchantDetail(merchantID string, merchantDetail *entity.MerchantDetail) error
}

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) DomainRepository {
	return &orderRepository{
		DB: DB,
	}
}

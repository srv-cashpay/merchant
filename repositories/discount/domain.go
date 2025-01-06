package discount

import (
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.DiscountRequest) (dto.DiscountResponse, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetByIdRequest) (*dto.DiscountResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.DiscountUpdateRequest) (dto.DiscountUpdateResponse, error)
}

type discountRepository struct {
	DB *gorm.DB
}

func NewDiscountRepository(DB *gorm.DB) DomainRepository {
	return &discountRepository{
		DB: DB,
	}
}

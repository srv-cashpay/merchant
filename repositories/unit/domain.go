package unit

import (
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.UnitRequest) (dto.UnitResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetByIdRequest) (*dto.UnitResponse, error)
	Update(req dto.UnitUpdateRequest) (dto.UnitUpdateResponse, error)
}

type unitRepository struct {
	DB *gorm.DB
}

func NewUnitRepository(DB *gorm.DB) DomainRepository {
	return &unitRepository{
		DB: DB,
	}
}

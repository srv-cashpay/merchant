package category

import (
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.CategoryRequest) (dto.CategoryResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetByIdRequest) (*dto.CategoryResponse, error)
	Update(req dto.CategoryUpdateRequest) (dto.CategoryUpdateResponse, error)
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) DomainRepository {
	return &categoryRepository{
		DB: DB,
	}
}

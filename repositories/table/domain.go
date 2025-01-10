package table

import (
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.TableRequest) (dto.TableResponse, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetTableByIdRequest) (*dto.TableResponse, error)
	Delete(req dto.DeleteTableRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.TableUpdateRequest) (dto.TableUpdateResponse, error)
}

type tableRepository struct {
	DB *gorm.DB
}

func NewTableRepository(DB *gorm.DB) DomainRepository {
	return &tableRepository{
		DB: DB,
	}
}

package merk

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.MerkRequest) (dto.MerkResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetByIdRequest) (*dto.MerkResponse, error)
	Update(req dto.MerkUpdateRequest) (dto.MerkUpdateResponse, error)
	CheckMerchantDetail(merchantID string, merchantDetail *entity.MerchantDetail) error
}

type merkRepository struct {
	DB *gorm.DB
}

func NewMerkRepository(DB *gorm.DB) DomainRepository {
	return &merkRepository{
		DB: DB,
	}
}

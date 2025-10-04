package voucher

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.VoucherRequest) (dto.VoucherResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetByIdRequest) (*dto.VoucherResponse, error)
	GetVerifikasi(req dto.GetVerifikasi) (*dto.GetVerifikasiResponse, error)
	Update(req dto.VoucherUpdateRequest) (dto.VoucherUpdateResponse, error)
	CheckMerchantDetail(merchantID string, merchantDetail *entity.MerchantDetail) error
}

type voucherRepository struct {
	DB *gorm.DB
}

func NewVoucherRepository(DB *gorm.DB) DomainRepository {
	return &voucherRepository{
		DB: DB,
	}
}

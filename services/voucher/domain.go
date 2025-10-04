package voucher

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/voucher"
)

type VoucherService interface {
	Create(req dto.VoucherRequest) (dto.VoucherResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdRequest) (*dto.VoucherResponse, error)
	GetVerifikasi(req dto.GetVerifikasi) (*dto.GetVerifikasiResponse, error)
	Update(req dto.VoucherUpdateRequest) (dto.VoucherUpdateResponse, error)
}

type voucherService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewVoucherService(Repo r.DomainRepository, jwtS m.JWTService) VoucherService {
	return &voucherService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

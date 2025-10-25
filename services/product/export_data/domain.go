package export_data

import (
	"context"

	m "github.com/srv-cashpay/middlewares/middlewares"
	"github.com/xuri/excelize/v2"

	dto "github.com/srv-cashpay/merchant/dto"
	r "github.com/srv-cashpay/merchant/repositories/product/export_data"
)

type ExportService interface {
	ExportExcel(ctx context.Context, req dto.ExportFilter) (*excelize.File, error)
}

type exportService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewExportService(Repo r.DomainRepository, jwtS m.JWTService) ExportService {
	return &exportService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

package table

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/table"
)

type TableService interface {
	Create(req dto.TableRequest) (dto.TableResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetTableByIdRequest) (*dto.TableResponse, error)
	Delete(req dto.DeleteTableRequest) (dto.DeleteTableResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.TableUpdateRequest) (dto.TableUpdateResponse, error)
}

type tableService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewTableService(Repo r.DomainRepository, jwtS m.JWTService) TableService {
	return &tableService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

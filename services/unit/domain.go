package unit

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/unit"
)

type UnitService interface {
	Create(req dto.UnitRequest) (dto.UnitResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdRequest) (*dto.UnitResponse, error)
	Update(req dto.UnitUpdateRequest) (dto.UnitUpdateResponse, error)
}

type unitService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewUnitService(Repo r.DomainRepository, jwtS m.JWTService) UnitService {
	return &unitService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

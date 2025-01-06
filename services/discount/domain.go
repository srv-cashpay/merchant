package discount

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/discount"
)

type DiscountService interface {
	Create(req dto.DiscountRequest) (dto.DiscountResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdRequest) (*dto.DiscountResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.DiscountUpdateRequest) (dto.DiscountUpdateResponse, error)
}

type discountService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewDiscountService(Repo r.DomainRepository, jwtS m.JWTService) DiscountService {
	return &discountService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

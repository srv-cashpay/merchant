package category

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/category"
)

type CategoryService interface {
	Create(req dto.CategoryRequest) (dto.CategoryResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdRequest) (*dto.CategoryResponse, error)
	Update(req dto.CategoryUpdateRequest) (dto.CategoryUpdateResponse, error)
}

type categoryService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewCategoryService(Repo r.DomainRepository, jwtS m.JWTService) CategoryService {
	return &categoryService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

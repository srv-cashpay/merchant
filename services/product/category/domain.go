package dashboard

import (
	m "github.com/srv-cashpay/middlewares/middlewares"

	"github.com/srv-cashpay/merchant/dto"
	r "github.com/srv-cashpay/merchant/repositories/product/category"
)

type GetCategoryService interface {
	Get(req dto.CategoryRequest) ([]dto.CategoryResponse, error)
}

type getCategorydashboardService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewGetCategoryService(Repo r.DomainRepository, jwtS m.JWTService) GetCategoryService {
	return &getCategorydashboardService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

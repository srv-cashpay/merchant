package order

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/order"
)

type OrderService interface {
	Create(req dto.OrderRequest) (dto.OrderResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Delete(req dto.DeleteOrderRequest) (dto.DeleteOrderResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdOrderRequest) (*dto.OrderResponse, error)
	Update(req dto.OrderUpdateRequest) (dto.OrderUpdateResponse, error)
}

type orderService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewOrderService(Repo r.DomainRepository, jwtS m.JWTService) OrderService {
	return &orderService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

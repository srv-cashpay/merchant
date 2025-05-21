package history

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/subscribe/history"
)

type HistoryService interface {
	Get(context echo.Context, req *dto.Pagination) dto.Response
	CheckAndExpire(orderID string) (*entity.Subscribe, error)
	GetById(req dto.GetHistory) (*dto.VAResponse, error)
}

type historyService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewHistoryService(Repo r.DomainRepository, jwtS m.JWTService) HistoryService {
	return &historyService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

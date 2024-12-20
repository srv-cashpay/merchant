package dashboard

import (
	m "github.com/srv-cashpay/middlewares/middlewares"

	"github.com/srv-cashpay/merchant/dto"
	r "github.com/srv-cashpay/merchant/repositories/dashboard"
)

type DashboardService interface {
	Get(req dto.GetDashboardRequest) (dto.GetDashboardResponse, error)
}

type dashboardService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewDashboardService(Repo r.DomainRepository, jwtS m.JWTService) DashboardService {
	return &dashboardService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

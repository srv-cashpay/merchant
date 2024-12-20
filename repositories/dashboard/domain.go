package dashboard

import (
	"github.com/srv-cashpay/merchant/dto"
	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req dto.GetDashboardRequest) (dto.GetDashboardResponse, error)
}

type dashboardRepository struct {
	DB *gorm.DB
}

func NewDashboardRepository(DB *gorm.DB) DomainRepository {
	return &dashboardRepository{DB: DB}
}

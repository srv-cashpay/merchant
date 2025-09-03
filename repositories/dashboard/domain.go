package dashboard

import (
	"github.com/srv-cashpay/merchant/dto"
	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req dto.GetDashboardRequest) (dto.GetDashboardResponse, error)
	SaveToken(req dto.TokenRequest) error
	GetAllTokens() ([]string, error)

	DeleteToken(token string) error
}

type dashboardRepository struct {
	DB *gorm.DB
}

func NewDashboardRepository(DB *gorm.DB) DomainRepository {
	return &dashboardRepository{DB: DB}
}

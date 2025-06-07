package contentsetting

import (
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/dashboard/contentsetting"
)

type ContentSettingService interface {
	Get(req dto.ContentSettingRequest) (dto.ContentSettingResponse, error)
	Update(req dto.UpdateContentSettingRequest) (dto.UpdateContentSettingResponse, error)
}

type contentsettingService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewContentSettingService(Repo r.DomainRepository, jwtS m.JWTService) ContentSettingService {
	return &contentsettingService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

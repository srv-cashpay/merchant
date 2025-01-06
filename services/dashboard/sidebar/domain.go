package sidebar

import (
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/dashboard/sidebar"
)

type SidebarService interface {
	Create(req dto.SidebarRequest) (dto.SidebarResponse, error)
	Get(req dto.SidebarRequest) (dto.GetSidebarResponse, error)
	GetById(req dto.GetSidebarByIdRequest) (*dto.SidebarResponse, error)
	Delete(req dto.DeleteSidebarRequest) (dto.DeleteSidebarResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.SidebarUpdateRequest) (dto.SidebarUpdateResponse, error)
}

type sidebarService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewSidebarService(Repo r.DomainRepository, jwtS m.JWTService) SidebarService {
	return &sidebarService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

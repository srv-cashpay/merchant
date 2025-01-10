package permission

import (
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/dashboard/permission"
)

type PermissionService interface {
	Create(req dto.PermissionRequest) (dto.PermissionResponse, error)
	Get(req dto.PermissionRequest) (dto.GetPermissionResponse, error)
	GetById(req dto.GetPermissionByIdRequest) (*dto.PermissionResponse, error)
	Delete(req dto.DeletePermissionRequest) (dto.DeletePermissionResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.PermissionUpdateRequest) (dto.PermissionUpdateResponse, error)
}

type permissionService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPermissionService(Repo r.DomainRepository, jwtS m.JWTService) PermissionService {
	return &permissionService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

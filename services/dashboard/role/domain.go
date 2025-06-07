package role

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/dashboard/role"
)

type RoleService interface {
	Create(req dto.RoleRequest) (dto.RoleResponse, error)
	Get(req dto.RoleUserRequest) (dto.GetRoleResponse, error)
	Pagination(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetRoleByIdRequest) (*dto.RoleResponse, error)
	Delete(req dto.DeleteRoleRequest) (dto.DeleteRoleResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.RoleUpdateRequest) (dto.RoleUpdateResponse, error)
}

type roleService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewRoleService(Repo r.DomainRepository, jwtS m.JWTService) RoleService {
	return &roleService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

package merk

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/merk"
)

type MerkService interface {
	Create(req dto.MerkRequest) (dto.MerkResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdRequest) (*dto.MerkResponse, error)
	Update(req dto.MerkUpdateRequest) (dto.MerkUpdateResponse, error)
}

type merkService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewMerkService(Repo r.DomainRepository, jwtS m.JWTService) MerkService {
	return &merkService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

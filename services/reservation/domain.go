package reservation

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/merchant/repositories/reservation"
)

type ReservationService interface {
	Create(req dto.ReservationRequest) (dto.ReservationResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetReservationByIdRequest) (*dto.ReservationResponse, error)
	Delete(req dto.DeleteReservationRequest) (dto.DeleteReservationResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.UpdateReservationRequest) (dto.UpdateReservationResponse, error)
}

type reservationService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewReservationService(Repo r.DomainRepository, jwtS m.JWTService) ReservationService {
	return &reservationService{
		Repo: Repo,
		jwt:  jwtS,
	}
}

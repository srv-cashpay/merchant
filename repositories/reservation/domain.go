package reservation

import (
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.ReservationRequest) (dto.ReservationResponse, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetReservationByIdRequest) (*dto.ReservationResponse, error)
	Delete(req dto.DeleteReservationRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.UpdateReservationRequest) (dto.UpdateReservationResponse, error)
}

type reservationRepository struct {
	DB *gorm.DB
}

func NewReservationRepository(DB *gorm.DB) DomainRepository {
	return &reservationRepository{
		DB: DB,
	}
}

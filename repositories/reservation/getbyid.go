package reservation

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *reservationRepository) GetById(req dto.GetReservationByIdRequest) (*dto.ReservationResponse, error) {
	tr := entity.Reservation{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.ReservationResponse{
		Reservation: tr.Table,
	}

	return response, nil
}

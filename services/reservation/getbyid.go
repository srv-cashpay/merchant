package reservation

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *reservationService) GetById(req dto.GetReservationByIdRequest) (*dto.ReservationResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

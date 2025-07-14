package reservation

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *reservationService) Delete(req dto.DeleteReservationRequest) (dto.DeleteReservationResponse, error) {
	transactionBody := dto.DeleteReservationRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeleteReservationResponse{}, err
	}

	response := dto.DeleteReservationResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}

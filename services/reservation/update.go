package reservation

import "github.com/srv-cashpay/merchant/dto"

func (b *reservationService) Update(req dto.UpdateReservationRequest) (dto.UpdateReservationResponse, error) {
	request := dto.UpdateReservationResponse{
		Floor:       req.Floor,
		UpdatedBy:   req.UpdatedBy,
		UserID:      req.UserID,
		Description: req.Description,
	}

	reservation, err := b.Repo.Update(req)
	if err != nil {
		return reservation, err
	}

	response := dto.UpdateReservationResponse{
		Floor:       request.Floor,
		UpdatedBy:   request.UpdatedBy,
		UserID:      request.UserID,
		Description: request.Description,
	}

	return response, nil
}

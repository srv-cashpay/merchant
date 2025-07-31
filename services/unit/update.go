package unit

import "github.com/srv-cashpay/merchant/dto"

func (b *unitService) Update(req dto.UnitUpdateRequest) (dto.UnitUpdateResponse, error) {
	request := dto.UnitUpdateRequest{
		UnitName: req.UnitName,
		UserID:   req.UserID,
		Status:   req.Status,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.UnitUpdateResponse{
		UnitName: request.UnitName,
		UserID:   request.UserID,
		Status:   request.Status,
	}

	return response, nil
}

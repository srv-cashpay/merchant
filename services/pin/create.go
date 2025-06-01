package pin

import (
	dto "github.com/srv-cashpay/merchant/dto"
	util "github.com/srv-cashpay/util/s"
)

func (s *pinService) Create(req dto.PinRequest) (dto.PinResponse, error) {

	// Proses pembuatan data Pin
	create := dto.PinRequest{
		ID:         util.GenerateRandomString(),
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		Pin:        req.Pin,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.PinResponse{}, err
	}

	response := dto.PinResponse{
		ID:         created.ID,
		MerchantID: created.MerchantID,
		Pin:        created.Pin,
		UserID:     created.UserID,
		CreatedBy:  created.CreatedBy,
	}

	return response, nil
}

package unit

import (
	dto "github.com/srv-cashpay/merchant/dto"
	util "github.com/srv-cashpay/util/s"
)

func (s *unitService) Create(req dto.UnitRequest) (dto.UnitResponse, error) {
	create := dto.UnitRequest{
		ID:         util.GenerateRandomString(),
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		UnitName:   req.UnitName,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.UnitResponse{}, err
	}

	response := dto.UnitResponse{
		ID:         created.ID,
		MerchantID: created.MerchantID,
		UnitName:   created.UnitName,
		UserID:     created.UserID,
		CreatedBy:  created.CreatedBy,
	}

	return response, nil
}

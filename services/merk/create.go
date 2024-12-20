package merk

import (
	dto "github.com/srv-cashpay/merchant/dto"
	util "github.com/srv-cashpay/util/s"
)

func (s *merkService) Create(req dto.MerkRequest) (dto.MerkResponse, error) {
	create := dto.MerkRequest{
		ID:         util.GenerateRandomString(),
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		MerkName:   req.MerkName,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.MerkResponse{}, err
	}

	response := dto.MerkResponse{
		ID:         created.ID,
		MerchantID: created.MerchantID,
		MerkName:   created.MerkName,
		UserID:     created.UserID,
		CreatedBy:  created.CreatedBy,
	}

	return response, nil
}

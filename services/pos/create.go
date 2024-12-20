package pos

import (
	dto "github.com/srv-cashpay/merchant/dto"
	util "github.com/srv-cashpay/util/s"
)

func (s *posService) Create(req dto.PosRequest) (dto.PosResponse, error) {
	create := dto.PosRequest{
		ID:        util.GenerateRandomString(),
		UserID:    req.UserID,
		CreatedBy: req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.PosResponse{}, err
	}

	response := dto.PosResponse{
		ID:        created.ID,
		UserID:    created.UserID,
		CreatedBy: created.CreatedBy,
	}

	return response, nil
}

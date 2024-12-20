package authenticator_request

import (
	dto "github.com/srv-cashpay/merchant/dto"
	util "github.com/srv-cashpay/util/s"
)

func (s *authenticatorService) Create(req dto.AuthenticatorRequest) (dto.AuthenticatorResponse, error) {
	create := dto.AuthenticatorRequest{
		ID:        util.GenerateRandomString(),
		UserID:    req.UserID,
		CreatedBy: req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.AuthenticatorResponse{}, err
	}

	response := dto.AuthenticatorResponse{
		ID:        created.ID,
		UserID:    created.UserID,
		CreatedBy: created.CreatedBy,
	}

	return response, nil
}

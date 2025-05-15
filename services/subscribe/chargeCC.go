package subscribe

import (
	"errors"

	dto "github.com/srv-cashpay/merchant/dto"
)

// TokenizeCard adalah service untuk men-tokenisasi kartu kredit
func (s *subscribeService) TokenizeCard(req dto.TokenizeRequest) (*dto.TokenizeResponse, error) {
	resp, err := s.Repo.TokenizeCard(req)
	if err != nil {
		return nil, err
	}

	if resp.Status != "success" {
		return nil, errors.New(resp.StatusMessage)
	}

	return resp, nil
}

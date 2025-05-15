package subscribe

import (
	"errors"

	dto "github.com/srv-cashpay/merchant/dto"
)

// TokenizeCard adalah service untuk men-tokenisasi kartu kredit
func (s *subscribeService) TokenizeCard(req dto.TokenizeRequest) (*dto.TokenizeResponse, error) {

	if req.OrderID == "" || req.Amount <= 0 {
		return nil, errors.New("missing required fields: order_id or amount")
	}

	resp, err := s.Repo.TokenizeCard(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != "201" {
		return nil, errors.New(resp.StatusMessage)
	}

	return resp, nil
}

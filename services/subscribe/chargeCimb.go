package subscribe

import (
	"errors"

	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *subscribeService) ChargeCimb(req dto.ChargeRequest) (*dto.VAResponse, error) {
	if req.OrderID == "" || req.GrossAmount <= 0 {
		return nil, errors.New("missing required fields: order_id or amount")
	}

	resp, err := s.Repo.ChargeCimb(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != "201" {
		return nil, errors.New(resp.StatusMessage)
	}

	return resp, nil
}

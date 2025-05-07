package subscribe

import (
	"errors"

	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *subscribeService) ChargeBca(req dto.ChargeRequest) (*dto.VAResponse, error) {
	if req.OrderID == "" || req.Amount <= 0 {
		return nil, errors.New("missing required fields: order_id or amount")
	}

	// Call the repository method to charge the BCA account
	response, err := s.Repo.ChargeBca(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

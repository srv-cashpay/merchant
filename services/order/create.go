package order

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *orderService) Create(req dto.OrderRequest) (dto.OrderResponse, error) {

	// Proses pembuatan data Order
	create := dto.OrderRequest{
		ID:         req.ID,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		OrderName:  req.OrderName,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	response := dto.OrderResponse{
		ID:         created.ID,
		MerchantID: created.MerchantID,
		OrderName:  created.OrderName,
		UserID:     created.UserID,
		CreatedBy:  created.CreatedBy,
	}

	return response, nil
}

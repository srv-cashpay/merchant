package order

import (
	"encoding/json"

	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *orderService) Create(req dto.OrderRequest) (dto.OrderResponse, error) {
	// Convert Product slice to JSON string
	productJSON, err := json.Marshal(req.Product)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	// Buat copy request dengan string product (opsional kalau mau simpan)
	reqWithJSON := req
	reqWithJSON.ProductJSON = string(productJSON)

	// Kirim ke repo
	created, err := s.Repo.Create(reqWithJSON)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	return created, nil
}

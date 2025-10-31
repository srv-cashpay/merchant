package tax

import (
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *taxService) Create(req dto.TaxRequest) (dto.TaxResponse, error) {
	if req.Status != 1 && req.Status != 2 {
		return dto.TaxResponse{}, fmt.Errorf("invalid status: must be 1 (active) or 2 (inactive)")
	}

	create := dto.TaxRequest{
		Tax:           req.Tax,
		TaxPercentage: req.TaxPercentage,
		Description:   req.Description,
		Status:        req.Status,
		UserID:        req.UserID,
		MerchantID:    req.MerchantID,
		CreatedBy:     req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.TaxResponse{}, err
	}

	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	// Dapatkan string status berdasarkan nilai integer
	statusString, ok := statusMap[create.Status]
	if !ok {
		return dto.TaxResponse{}, fmt.Errorf("invalid status value in database")
	}

	response := dto.TaxResponse{
		ID:            created.ID,
		UserID:        created.UserID,
		Tax:           created.Tax,
		TaxPercentage: req.TaxPercentage,
		Description:   created.Description,
		Status:        statusString,
		MerchantID:    created.MerchantID,
		CreatedBy:     created.CreatedBy,
	}

	return response, nil
}

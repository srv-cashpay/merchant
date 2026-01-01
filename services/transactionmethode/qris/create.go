package qris

import (
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *qrisService) Create(req dto.CoQrisRequest) (dto.CoQrisResponse, error) {
	if req.Status != 1 && req.Status != 2 {
		return dto.CoQrisResponse{}, fmt.Errorf("invalid status: must be 1 (active) or 2 (inactive)")
	}

	create := dto.CoQrisRequest{
		QrisName:   req.QrisName,
		Link:       req.Link,
		File:       req.File,
		Status:     req.Status,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.CoQrisResponse{}, err
	}

	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	// Dapatkan string status berdasarkan nilai integer
	statusString, ok := statusMap[create.Status]
	if !ok {
		return dto.CoQrisResponse{}, fmt.Errorf("invalid status value in database")
	}

	response := dto.CoQrisResponse{
		QrisName:   created.QrisName,
		Link:       created.Link,
		FilePath:   created.FilePath,
		Status:     statusString,
		UserID:     created.UserID,
		MerchantID: created.MerchantID,
		CreatedBy:  created.CreatedBy,
	}

	return response, nil
}

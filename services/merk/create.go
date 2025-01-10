package merk

import (
	"errors"
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	util "github.com/srv-cashpay/util/s"
	"gorm.io/gorm"
)

func (s *merkService) Create(req dto.MerkRequest) (dto.MerkResponse, error) {
	// Validasi MerchantDetail
	var merchantDetail entity.MerchantDetail
	err := s.Repo.CheckMerchantDetail(req.MerchantID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.MerkResponse{}, fmt.Errorf("merchant detail not found for merchant_id: %s", req.MerchantID)
		}
		return dto.MerkResponse{}, err
	}

	// Proses pembuatan data Merk
	create := dto.MerkRequest{
		ID:         util.GenerateRandomString(),
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		MerkName:   req.MerkName,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.MerkResponse{}, err
	}

	response := dto.MerkResponse{
		ID:         created.ID,
		MerchantID: created.MerchantID,
		MerkName:   created.MerkName,
		UserID:     created.UserID,
		CreatedBy:  created.CreatedBy,
	}

	return response, nil
}

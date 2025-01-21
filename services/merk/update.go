package merk

import (
	"errors"
	"fmt"

	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	"gorm.io/gorm"
)

func (b *merkService) Update(req dto.MerkUpdateRequest) (dto.MerkUpdateResponse, error) {
	var merchantDetail entity.MerchantDetail
	err := b.Repo.CheckMerchantDetail(req.MerchantID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.MerkUpdateResponse{}, fmt.Errorf("merchant detail not found for merchant_id: %s", req.MerchantID)
		}
		return dto.MerkUpdateResponse{}, err
	}
	request := dto.MerkUpdateRequest{
		MerkName:    req.MerkName,
		UserID:      req.UserID,
		MerchantID:  req.MerchantID,
		UpdatedBy:   req.UpdatedBy,
		Description: req.Description,
		Status:      req.Status,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.MerkUpdateResponse{
		MerkName:    request.MerkName,
		UserID:      request.UserID,
		MerchantID:  request.MerchantID,
		UpdatedBy:   request.UpdatedBy,
		Description: request.Description,
		Status:      request.Status,
	}

	return response, nil
}

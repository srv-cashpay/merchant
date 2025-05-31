package pin

import (
	"errors"
	"fmt"

	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	"gorm.io/gorm"
)

func (b *pinService) Update(req dto.PinUpdateRequest) (dto.PinUpdateResponse, error) {
	var merchantDetail entity.MerchantDetail
	err := b.Repo.CheckMerchantDetail(req.MerchantID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.PinUpdateResponse{}, fmt.Errorf("merchant detail not found for merchant_id: %s", req.MerchantID)
		}
		return dto.PinUpdateResponse{}, err
	}
	request := dto.PinUpdateRequest{
		Pin:         req.Pin,
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

	response := dto.PinUpdateResponse{
		Pin:         request.Pin,
		UserID:      request.UserID,
		MerchantID:  request.MerchantID,
		UpdatedBy:   request.UpdatedBy,
		Description: request.Description,
		Status:      request.Status,
	}

	return response, nil
}

package order

import (
	"errors"
	"fmt"

	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	"gorm.io/gorm"
)

func (b *orderService) Update(req dto.OrderUpdateRequest) (dto.OrderUpdateResponse, error) {
	var merchantDetail entity.MerchantDetail
	err := b.Repo.CheckMerchantDetail(req.MerchantID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.OrderUpdateResponse{}, fmt.Errorf("merchant detail not found for merchant_id: %s", req.MerchantID)
		}
		return dto.OrderUpdateResponse{}, err
	}
	request := dto.OrderUpdateRequest{
		OrderName:   req.OrderName,
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

	response := dto.OrderUpdateResponse{
		OrderName:   request.OrderName,
		UserID:      request.UserID,
		MerchantID:  request.MerchantID,
		UpdatedBy:   request.UpdatedBy,
		Description: request.Description,
		Status:      request.Status,
	}

	return response, nil
}

package category

import (
	"errors"
	"fmt"

	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	"gorm.io/gorm"
)

func (b *categoryService) Update(req dto.CategoryUpdateRequest) (dto.CategoryUpdateResponse, error) {
	var merchantDetail entity.MerchantDetail
	err := b.Repo.CheckMerchantDetail(req.MerchantID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.CategoryUpdateResponse{}, fmt.Errorf("merchant detail not found for merchant_id: %s", req.MerchantID)
		}
		return dto.CategoryUpdateResponse{}, err
	}
	request := dto.CategoryUpdateRequest{
		CategoryName: req.CategoryName,
		UserID:       req.UserID,
		Status:       req.Status,
		Description:  req.Description,
		UpdatedBy:    req.UpdatedBy,
		MerchantID:   req.MerchantID,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.CategoryUpdateResponse{
		CategoryName: request.CategoryName,
		UserID:       request.UserID,
		Status:       request.Status,
		Description:  request.Description,
		UpdatedBy:    request.UpdatedBy,
		MerchantID:   request.MerchantID,
	}

	return response, nil
}

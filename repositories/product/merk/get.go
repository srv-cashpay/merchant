package merk

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *getmerkRepository) Get(req dto.MerkRequest) ([]dto.MerkResponse, error) {
	var data []entity.Merk

	if err := r.DB.Where("status = ? AND merchant_id = ?", 1, req.MerchantID).Find(&data).Error; err != nil {
		return nil, err
	}

	var responses []dto.MerkResponse
	for _, category := range data {
		responses = append(responses, dto.MerkResponse{
			ID:          category.ID,
			UserID:      category.UserID,
			MerchantID:  category.MerchantID,
			Description: category.Description,
			MerkName:    category.MerkName,
			CreatedBy:   category.CreatedBy,
			Status:      category.Status,
		})
	}

	return responses, nil
}

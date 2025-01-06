package category

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *getcategoryRepository) Get(req dto.CategoryRequest) ([]dto.CategoryResponse, error) {
	var data []entity.Category

	if err := r.DB.Where("status = ? AND merchant_id = ?", 1, req.MerchantID).Find(&data).Error; err != nil {
		return nil, err
	}

	var responses []dto.CategoryResponse
	for _, category := range data {
		responses = append(responses, dto.CategoryResponse{
			ID:           category.ID,
			UserID:       category.UserID,
			MerchantID:   category.MerchantID,
			Description:  category.Description,
			CategoryName: category.CategoryName,
			CreatedBy:    category.CreatedBy,
			Status:       category.Status,
		})
	}

	return responses, nil
}

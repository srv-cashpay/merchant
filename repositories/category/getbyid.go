package category

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *categoryRepository) GetById(req dto.GetByIdRequest) (*dto.CategoryResponse, error) {
	tr := entity.Category{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.CategoryResponse{
		CategoryName: tr.CategoryName,
	}

	return response, nil
}

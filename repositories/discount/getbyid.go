package discount

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *discountRepository) GetById(req dto.GetByIdRequest) (*dto.DiscountResponse, error) {
	tr := entity.Discount{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.DiscountResponse{
		DiscountName:       tr.DiscountName,
		DiscountPercentage: tr.DiscountPercentage,
	}

	return response, nil
}

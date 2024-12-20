package category

import (
	"github.com/srv-cashpay/merchant/entity"
)

func (r *getcategoryRepository) Get() ([]entity.Category, error) {
	var data []entity.Category

	if err := r.DB.Where("status = ?", 1).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

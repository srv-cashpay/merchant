package merk

import (
	"github.com/srv-cashpay/merchant/entity"
)

func (r *getmerkRepository) Get() ([]entity.Merk, error) {
	var data []entity.Merk

	if err := r.DB.Where("status = ?", 1).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

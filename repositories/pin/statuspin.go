package pin

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *pinRepository) GetPinStatus(req dto.PinRequest) (*entity.Pin, error) {
	var pin entity.Pin
	if err := r.DB.Where("user_id = ?", req.UserID).First(&pin).Error; err != nil {
		return nil, err
	}
	return &pin, nil
}

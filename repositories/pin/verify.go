package pin

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *pinRepository) Verify(req dto.VerifyPinRequest) (*entity.Pin, error) {
	var pin entity.Pin

	err := r.DB.Where("user_id = ?", req.UserID).First(&pin).Error
	if err != nil {
		return nil, err
	}

	return &pin, nil
}

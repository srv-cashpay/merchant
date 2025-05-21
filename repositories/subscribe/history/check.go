package history

import (
	"time"

	"github.com/srv-cashpay/merchant/entity"
)

func (r *historyRepository) CheckAndExpireIfNeeded(orderID string) (*entity.Subscribe, error) {
	var sub entity.Subscribe

	if err := r.DB.Where("order_id = ?", orderID).First(&sub).Error; err != nil {
		return nil, err
	}

	if sub.Status == "pending" && time.Now().After(sub.ExpiryTime) {
		sub.Status = "expired"
		if err := r.DB.Save(&sub).Error; err != nil {
			return nil, err
		}
	}

	return &sub, nil
}

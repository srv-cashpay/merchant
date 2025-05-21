package history

import "github.com/srv-cashpay/merchant/entity"

func (s *historyService) CheckAndExpire(orderID string) (*entity.Subscribe, error) {
	return s.Repo.CheckAndExpireIfNeeded(orderID)
}

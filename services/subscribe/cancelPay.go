package subscribe

import "github.com/srv-cashpay/merchant/dto"

func (s *subscribeService) CancelPay(req dto.GetorderID) ([]byte, int, error) {
	return s.Repo.CancelPay(req)
}

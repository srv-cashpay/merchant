package dashboard

import "github.com/srv-cashpay/merchant/dto"

func (s *dashboardService) Get(req dto.GetDashboardRequest) (dto.GetDashboardResponse, error) {
	products, _ := s.Repo.Get(req)

	return products, nil
}

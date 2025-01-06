package sidebar

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *sidebarService) Get(req dto.SidebarRequest) (dto.GetSidebarResponse, error) {
	products, _ := s.Repo.Get(req)

	return products, nil
}

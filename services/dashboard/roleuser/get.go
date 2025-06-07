package roleuser

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *roleuserService) Get(req dto.RoleUserRequest) (dto.GetRoleUserResponse, error) {
	products, _ := s.Repo.Get(req)

	return products, nil
}

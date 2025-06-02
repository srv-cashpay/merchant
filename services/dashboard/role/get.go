package role

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *roleService) Get(req dto.RoleUser) (dto.GetRoleResponse, error) {
	products, _ := s.Repo.Get(req)

	return products, nil
}

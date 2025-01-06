package sidebar

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *sidebarService) GetById(req dto.GetSidebarByIdRequest) (*dto.SidebarResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

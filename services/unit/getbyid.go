package unit

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *unitService) GetById(req dto.GetByIdRequest) (*dto.UnitResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

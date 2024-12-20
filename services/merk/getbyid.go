package merk

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *merkService) GetById(req dto.GetByIdRequest) (*dto.MerkResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

package table

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *tableService) GetById(req dto.GetTableByIdRequest) (*dto.TableResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

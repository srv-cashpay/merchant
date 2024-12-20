package category

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *categoryService) GetById(req dto.GetByIdRequest) (*dto.CategoryResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

package table

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *tableService) Delete(req dto.DeleteTableRequest) (dto.DeleteTableResponse, error) {
	transactionBody := dto.DeleteTableRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeleteTableResponse{}, err
	}

	response := dto.DeleteTableResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}

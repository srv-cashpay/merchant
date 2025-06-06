package user

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *userService) Delete(req dto.DeleteRequest) (dto.DeleteResponse, error) {
	transactionBody := dto.DeleteRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeleteResponse{}, err
	}

	response := dto.DeleteResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}

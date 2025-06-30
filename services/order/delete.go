package order

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *orderService) Delete(req dto.DeleteOrderRequest) (dto.DeleteOrderResponse, error) {
	transactionBody := dto.DeleteOrderRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeleteOrderResponse{}, err
	}

	response := dto.DeleteOrderResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}

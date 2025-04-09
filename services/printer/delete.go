package printer

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *printerService) Delete(req dto.DeletePrinterRequest) (dto.DeletePrinterResponse, error) {
	transactionBody := dto.DeleteRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeletePrinterResponse{}, err
	}

	response := dto.DeletePrinterResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}

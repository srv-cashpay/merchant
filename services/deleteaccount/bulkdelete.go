package deleteaccount

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *deleteaccountService) BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error) {
	transactionBody := dto.BulkDeleteRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	count, err := b.Repo.BulkDelete(req)
	if err != nil {
		return dto.BulkDeleteResponse{}, err
	}

	response := dto.BulkDeleteResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
		Count:     count, // Menyimpan jumlah yang dihapus
	}

	return response, nil
}

package sidebar

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *sidebarService) Delete(req dto.DeleteSidebarRequest) (dto.DeleteSidebarResponse, error) {
	transactionBody := dto.DeleteSidebarRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeleteSidebarResponse{}, err
	}

	response := dto.DeleteSidebarResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}

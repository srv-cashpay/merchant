package merk

import "github.com/srv-cashpay/merchant/dto"

func (b *merkService) Update(req dto.MerkUpdateRequest) (dto.MerkUpdateResponse, error) {
	request := dto.MerkUpdateRequest{
		MerkName:    req.MerkName,
		UserID:      req.UserID,
		MerchantID:  req.MerchantID,
		UpdatedBy:   req.UpdatedBy,
		Description: req.Description,
		Status:      req.Status,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.MerkUpdateResponse{
		MerkName:    request.MerkName,
		UserID:      request.UserID,
		MerchantID:  request.MerchantID,
		UpdatedBy:   request.UpdatedBy,
		Description: request.Description,
		Status:      request.Status,
	}

	return response, nil
}

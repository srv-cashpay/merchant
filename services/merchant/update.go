package product

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *merchantService) Update(req dto.UpdateMerchantRequest) (dto.UpdateMerchantResponse, error) {
	request := dto.UpdateMerchantRequest{
		ID:           req.ID,
		MerchantName: req.MerchantName,
		IDNumber:     req.IDNumber,
		Description:  req.Description,
		Address:      req.Address,
		Country:      req.Country,
		CurrencyID:   req.CurrencyID,
		City:         req.City,
		Zip:          req.Zip,
		Phone:        req.Phone,
		UpdatedBy:    req.UpdatedBy,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.UpdateMerchantResponse{
		ID:           req.ID,
		MerchantName: request.MerchantName,
		IDNumber:     request.IDNumber,
		Description:  request.Description,
		Address:      request.Address,
		Country:      request.Country,
		CurrencyID:   request.CurrencyID,
		City:         request.City,
		Zip:          request.Zip,
		Phone:        request.Phone,
		UpdatedBy:    request.UpdatedBy,
	}

	return response, nil
}

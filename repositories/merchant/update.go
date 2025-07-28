package merchant

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *merchantRepository) Update(req dto.UpdateMerchantRequest) (dto.UpdateMerchantResponse, error) {
	tr := dto.GetByIdRequest{
		ID: req.ID,
	}

	request := entity.MerchantDetail{
		ID:           tr.ID,
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

	mer, err := b.GetById(tr)
	if err != nil {
		return dto.UpdateMerchantResponse{}, err
	}

	err = b.DB.Where("ID = ?", req.ID).Updates(entity.MerchantDetail{
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
	}).Error
	if err != nil {
		return dto.UpdateMerchantResponse{}, err
	}

	response := dto.UpdateMerchantResponse{
		MerchantName: request.MerchantName,
		Description:  request.Description,
		Address:      request.Address,
		Country:      request.Country,
		CurrencyID:   request.CurrencyID,
		City:         request.City,
		Zip:          request.Zip,
		Phone:        request.Phone,
		UpdatedBy:    request.UpdatedBy,
		ID:           mer.ID,
	}

	return response, nil
}

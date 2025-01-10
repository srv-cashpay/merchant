package merchant

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *merchantRepository) GetById(req dto.GetByIdRequest) (*dto.GetMerchantResponse, error) {
	tr := entity.MerchantDetail{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.GetMerchantResponse{
		ID:           tr.ID,
		MerchantName: tr.MerchantName,
		Description:  tr.Description,
		Address:      tr.Address,
		City:         tr.City,
		Country:      tr.Country,
		Zip:          tr.Zip,
		Phone:        tr.Phone,
		UpdatedBy:    tr.UpdatedBy,
	}

	return response, nil
}

package merchant

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *merchantRepository) Get(req dto.GetMerchantRequest) (dto.GetMerchantResponse, error) {

	var data entity.MerchantDetail

	if err := r.DB.Where("user_id = ?", req.UserID).Find(&data).Error; err != nil {
		return dto.GetMerchantResponse{}, err
	}

	response := dto.GetMerchantResponse{
		ID:           data.ID,
		UserID:       data.UserID,
		MerchantName: data.MerchantName,
		Address:      data.Address,
		City:         data.City,
		Zip:          data.Zip,
		Country:      data.Country,
		Phone:        data.Phone,
		CurrencyID:   data.CurrencyID,
		Description:  data.Description,
		UpdatedBy:    data.UpdatedBy,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		DeletedAt:    data.DeletedAt,
	}

	return response, nil
}

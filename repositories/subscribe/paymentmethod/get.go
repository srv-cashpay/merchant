package paymentmethod

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *paymentmethodRepository) Get(req dto.PaymentMethodRequest) ([]dto.PaymentMethodResponse, error) {
	var paymentmethods []entity.PaymentMethod
	if err := r.DB.Find(&paymentmethods).Error; err != nil {
		return nil, err
	}

	var result []dto.PaymentMethodResponse
	for _, pm := range paymentmethods {
		result = append(result, dto.PaymentMethodResponse{
			ID:            pm.ID,
			UserID:        pm.UserID,
			MerchantID:    pm.MerchantID,
			PaymentMethod: pm.PaymentMethod,
			Status:        pm.Status,
			CreatedBy:     pm.CreatedBy,
			UpdatedBy:     pm.UpdatedBy,
			DeletedBy:     pm.DeletedBy,
			CreatedAt:     pm.CreatedAt,
		})
	}

	return result, nil
}

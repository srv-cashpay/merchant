package paymentmethod

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *paymentmethodRepository) Get(req dto.PaymentMethodRequest) (dto.PaymentMethodResponse, error) {
	var paymentmethods []entity.PaymentMethod
	if err := r.DB.Preload("Image").Find(&paymentmethods).Error; err != nil {
		return dto.PaymentMethodResponse{}, err
	}

	// Bungkus izin dalam field 'items' seperti yang diharapkan
	return dto.PaymentMethodResponse{
		PaymentMethod: req.PaymentMethod,
	}, nil
}

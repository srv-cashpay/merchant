package payment

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *paymentRepository) GetById(req dto.GetByIdRequest) (*dto.PaymentResponse, error) {
	tr := entity.Payment{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.PaymentResponse{
		PaymentName:       tr.PaymentName,
		PaymentPercentage: tr.PaymentPercentage,
	}

	return response, nil
}

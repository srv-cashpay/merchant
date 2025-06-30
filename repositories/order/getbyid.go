package order

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *orderRepository) GetById(req dto.GetByIdOrderRequest) (*dto.OrderResponse, error) {
	tr := entity.Order{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.OrderResponse{
		OrderName: tr.OrderName,
		Status:    tr.Status,
	}

	return response, nil
}

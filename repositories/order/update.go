package order

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *orderRepository) Update(req dto.OrderUpdateRequest) (dto.OrderUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateOrder := entity.Order{
		OrderName:  req.OrderName,
		UpdatedBy:  req.UpdatedBy,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		Status:     req.Status,
	}

	var existingProduct entity.Order
	err := b.DB.Where("id = ?", req.ID).First(&existingProduct).Error
	if err != nil {
		return dto.OrderUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingProduct).Updates(updateOrder).Error
	if err != nil {
		return dto.OrderUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.OrderUpdateResponse{
		OrderName:  updateOrder.OrderName,
		UpdatedBy:  updateOrder.UpdatedBy,
		UserID:     updateOrder.UserID,
		MerchantID: updateOrder.MerchantID,
		Status:     updateOrder.Status,
	}

	return response, nil
}

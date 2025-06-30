package order

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *orderRepository) Order(req dto.OrderRequest) (dto.OrderResponse, error) {
	// Insert or update the auto_increment value based on merchant_id

	create := entity.Order{
		ID:         req.ID,
		OrderName:  req.OrderName,
		MerchantID: req.MerchantID,
		UserID:     req.UserID,
		CreatedBy:  req.CreatedBy,
		Product:    req.ProductJSON, // simpan JSON string
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.OrderResponse{}, err
	}

	response := dto.OrderResponse{
		ID:         req.ID,
		OrderName:  create.OrderName,
		MerchantID: req.MerchantID,
		UserID:     req.UserID,
		CreatedBy:  req.CreatedBy,
		Product:    req.Product, // simpan JSON string
	}

	return response, nil

}

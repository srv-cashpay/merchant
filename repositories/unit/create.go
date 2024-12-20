package unit

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *unitRepository) Create(req dto.UnitRequest) (dto.UnitResponse, error) {

	create := entity.Unit{
		ID:         req.ID,
		UnitName:   req.UnitName,
		MerchantID: req.MerchantID,
		UserID:     req.UserID,
		CreatedBy:  req.CreatedBy,
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.UnitResponse{}, err
	}

	response := dto.UnitResponse{
		ID:         req.ID,
		UnitName:   create.UnitName,
		MerchantID: req.MerchantID,
		UserID:     req.UserID,
		CreatedBy:  req.CreatedBy,
	}

	return response, nil

}

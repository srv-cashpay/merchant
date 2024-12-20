package pos

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/pos/entity"
)

func (r *posRepository) Create(req dto.PosRequest) (dto.PosResponse, error) {

	create := entity.Pos{
		ID:        req.ID,
		UserID:    req.UserID,
		CreatedBy: req.CreatedBy,
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.PosResponse{}, err
	}

	response := dto.PosResponse{
		ID:        req.ID,
		UserID:    req.UserID,
		CreatedBy: req.CreatedBy,
	}

	return response, nil

}

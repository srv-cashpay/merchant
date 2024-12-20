package authenticator_request

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/pos/entity"
)

func (r *authenticatorRepository) Create(req dto.AuthenticatorRequest) (dto.AuthenticatorResponse, error) {

	create := entity.Pos{
		ID:        req.ID,
		UserID:    req.UserID,
		CreatedBy: req.CreatedBy,
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.AuthenticatorResponse{}, err
	}

	response := dto.AuthenticatorResponse{
		ID:        req.ID,
		UserID:    req.UserID,
		CreatedBy: req.CreatedBy,
	}

	return response, nil

}

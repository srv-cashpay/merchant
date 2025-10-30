package user

import (
	"github.com/srv-cashpay/auth/entity"
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *userRepository) GetById(req dto.GetByIdRequest) (*dto.UserMerchantResponse, error) {
	tr := entity.AccessDoor{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.UserMerchantResponse{
		FullName: tr.FullName,
	}

	return response, nil
}

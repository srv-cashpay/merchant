package unit

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *unitRepository) GetById(req dto.GetByIdRequest) (*dto.UnitResponse, error) {
	tr := entity.Unit{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.UnitResponse{
		UnitName: tr.UnitName,
		Status:   tr.Status,
	}

	return response, nil
}

package merk

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *merkRepository) GetById(req dto.GetByIdRequest) (*dto.MerkResponse, error) {
	tr := entity.Merk{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.MerkResponse{
		MerkName: tr.MerkName,
	}

	return response, nil
}

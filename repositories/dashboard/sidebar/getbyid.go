package Sidebar

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *SidebarRepository) GetById(req dto.GetSidebarByIdRequest) (*dto.SidebarResponse, error) {
	tr := entity.Sidebar{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.SidebarResponse{
		Label: tr.Label,
	}

	return response, nil
}

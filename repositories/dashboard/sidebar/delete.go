package Sidebar

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *SidebarRepository) Delete(req dto.DeleteSidebarRequest) (dto.DeleteResponse, error) {
	tr := dto.GetSidebarByIdRequest{
		ID: req.ID,
	}

	_, err := b.GetById(tr)
	if err != nil {
		return dto.DeleteResponse{}, err
	}

	// Use GORM BeforeDelete hook to set DeletedBy
	if err := b.DB.Where("id = ?", req.ID).Delete(&entity.Sidebar{}).Error; err != nil {
		return dto.DeleteResponse{}, err
	}

	response := dto.DeleteResponse{
		DeletedBy: req.DeletedBy,
	}

	return response, nil
}

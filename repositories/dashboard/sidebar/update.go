package Sidebar

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *SidebarRepository) Update(req dto.SidebarUpdateRequest) (dto.SidebarUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateSidebar := entity.Sidebar{
		Label:      req.Label,
		Icon:       req.Icon, // Pastikan status boolean diterima dengan benar
		UpdatedBy:  req.UpdatedBy,
		UserID:     req.UserID,
		To:         req.To,
		MerchantID: req.MerchantID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingSidebar entity.Sidebar
	err := b.DB.Where("id = ?", req.ID).First(&existingSidebar).Error
	if err != nil {
		return dto.SidebarUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingSidebar).Updates(updateSidebar).Error
	if err != nil {
		return dto.SidebarUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.SidebarUpdateResponse{
		Label:      updateSidebar.Label,
		Icon:       updateSidebar.Icon,
		UpdatedBy:  updateSidebar.UpdatedBy,
		UserID:     updateSidebar.UserID,
		MerchantID: updateSidebar.MerchantID,
		To:         updateSidebar.To,
	}

	return response, nil
}

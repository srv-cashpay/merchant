package Sidebar

import (
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *SidebarRepository) Get(req dto.SidebarRequest) (dto.GetSidebarResponse, error) {
	var getdb []entity.Sidebar

	// Query menggunakan GORM
	err := r.DB.Where("merchant_id = ? AND user_id = ?", req.MerchantID, req.UserID).Find(&getdb).Error
	if err != nil {
		return dto.GetSidebarResponse{}, err
	}

	// Validasi jika data tidak ditemukan
	if len(getdb) == 0 {
		return dto.GetSidebarResponse{}, fmt.Errorf("no sidebar data found for merchant_id: %s and user_id: %s", req.MerchantID, req.UserID)
	}

	// Mengonversi data entity ke DTO
	var items []dto.SidebarItem
	for _, item := range getdb {
		items = append(items, dto.SidebarItem{
			Label: item.Label, // Sesuaikan dengan struktur data
			Icon:  item.Icon,
			To:    item.To,
		})
	}

	response := dto.GetSidebarResponse{
		Items: items,
	}

	return response, nil
}

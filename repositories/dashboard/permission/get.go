package permission

import (
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *PermissionRepository) Get(req dto.PermissionRequest) (dto.GetPermissionResponse, error) {
	var getdb []entity.Permission

	// Query menggunakan GORM
	err := r.DB.Where("merchant_id = ? AND user_id = ?", req.MerchantID, req.UserID).Find(&getdb).Error
	if err != nil {
		return dto.GetPermissionResponse{}, err
	}

	// Validasi jika data tidak ditemukan
	if len(getdb) == 0 {
		return dto.GetPermissionResponse{}, fmt.Errorf("no sidebar data found for merchant_id: %s and user_id: %s", req.MerchantID, req.UserID)
	}

	// Mengonversi data entity ke DTO
	var items []dto.PermissionItem
	for _, item := range getdb {
		items = append(items, dto.PermissionItem{
			Label: item.Label, // Sesuaikan dengan struktur data
			Icon:  item.Icon,
			To:    item.To,
		})
	}

	response := dto.GetPermissionResponse{
		Items: items,
	}

	return response, nil
}

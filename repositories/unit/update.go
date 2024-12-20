package unit

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *unitRepository) Update(req dto.UnitUpdateRequest) (dto.UnitUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateUnit := entity.Unit{
		UnitName:  req.UnitName,
		UpdatedBy: req.UpdatedBy,
		UserID:    req.UserID,
	}

	var existingProduct entity.Unit
	err := b.DB.Where("user_id = ?", req.UserID).First(&existingProduct).Error
	if err != nil {
		return dto.UnitUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingProduct).Updates(updateUnit).Error
	if err != nil {
		return dto.UnitUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.UnitUpdateResponse{
		UnitName:  updateUnit.UnitName,
		UpdatedBy: updateUnit.UpdatedBy,
		UserID:    updateUnit.UserID,
	}

	return response, nil
}

package merk

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *merkRepository) Update(req dto.MerkUpdateRequest) (dto.MerkUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateMerk := entity.Merk{
		MerkName:  req.MerkName,
		UpdatedBy: req.UpdatedBy,
		UserID:    req.UserID,
	}

	var existingProduct entity.Merk
	err := b.DB.Where("user_id = ?", req.UserID).First(&existingProduct).Error
	if err != nil {
		return dto.MerkUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingProduct).Updates(updateMerk).Error
	if err != nil {
		return dto.MerkUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.MerkUpdateResponse{
		MerkName:  updateMerk.MerkName,
		UpdatedBy: updateMerk.UpdatedBy,
		UserID:    updateMerk.UserID,
	}

	return response, nil
}

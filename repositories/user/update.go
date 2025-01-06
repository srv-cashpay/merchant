package user

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *userRepository) Update(req dto.UserUpdateRequest) (dto.UserUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateUser := entity.User{
		User:        req.User,
		Status:      req.Status, // Pastikan status boolean diterima dengan benar
		UpdatedBy:   req.UpdatedBy,
		UserID:      req.UserID,
		Description: req.Description,
		MerchantID:  req.MerchantID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingUser entity.User
	err := b.DB.Where("id = ?", req.ID).First(&existingUser).Error
	if err != nil {
		return dto.UserUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingUser).Updates(updateUser).Error
	if err != nil {
		return dto.UserUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.UserUpdateResponse{
		User:        updateUser.User,
		Status:      updateUser.Status,
		UpdatedBy:   updateUser.UpdatedBy,
		UserID:      updateUser.UserID,
		MerchantID:  updateUser.MerchantID,
		Description: updateUser.Description,
	}

	return response, nil
}

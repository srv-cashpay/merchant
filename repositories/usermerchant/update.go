package user

import (
	"github.com/srv-cashpay/auth/entity"
	"github.com/srv-cashpay/merchant/dto"
)

func (b *userRepository) Update(req dto.UserMerchantUpdateRequest) (dto.UserMerchantUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateUserMerchant := entity.AccessDoor{
		FullName: req.FullName,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingUserMerchant entity.AccessDoor
	err := b.DB.Where("id = ?", req.ID).First(&existingUserMerchant).Error
	if err != nil {
		return dto.UserMerchantUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingUserMerchant).Updates(updateUserMerchant).Error
	if err != nil {
		return dto.UserMerchantUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.UserMerchantUpdateResponse{
		FullName: updateUserMerchant.FullName,
	}

	return response, nil
}

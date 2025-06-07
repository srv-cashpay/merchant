package roleuser

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *RoleUserRepository) Update(req dto.RoleUserUpdateRequest) (dto.RoleUserUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateRole := entity.RoleUser{
		RoleID: req.Role,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingRole entity.RoleUser
	err := b.DB.Where("id = ?", req.ID).First(&existingRole).Error
	if err != nil {
		return dto.RoleUserUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingRole).Updates(updateRole).Error
	if err != nil {
		return dto.RoleUserUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.RoleUserUpdateResponse{
		Role: updateRole.RoleID,
	}

	return response, nil
}

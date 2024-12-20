package category

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *categoryRepository) Update(req dto.CategoryUpdateRequest) (dto.CategoryUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateCategory := entity.Category{
		CategoryName: req.CategoryName,
		UpdatedBy:    req.UpdatedBy,
		UserID:       req.UserID,
		Status:       req.Status,
		Description:  req.Description,
		MerchantID:   req.MerchantID,
	}

	var existingProduct entity.Category
	err := b.DB.Where("user_id = ?", req.UserID).First(&existingProduct).Error
	if err != nil {
		return dto.CategoryUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingProduct).Updates(updateCategory).Error
	if err != nil {
		return dto.CategoryUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.CategoryUpdateResponse{
		CategoryName: updateCategory.CategoryName,
		UpdatedBy:    updateCategory.UpdatedBy,
		UserID:       updateCategory.UserID,
		Status:       updateCategory.Status,
		Description:  updateCategory.Description,
	}

	return response, nil
}

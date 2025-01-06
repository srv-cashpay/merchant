package category

import (
	"fmt"
	"time"

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
		UpdatedAt:    time.Now(),
	}
	fmt.Println("Waktu update:", updateCategory.UpdatedAt)

	var existingProduct entity.Category
	err := b.DB.Where("id = ?", req.ID).First(&existingProduct).Error
	if err != nil {
		return dto.CategoryUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingProduct).Updates(updateCategory).Error
	if err != nil {
		return dto.CategoryUpdateResponse{}, err
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	waktuSekarang := time.Now().In(loc).Format("02 Januari 2006, 15:04:05")
	// Menyiapkan response setelah pembaruan berhasil
	response := dto.CategoryUpdateResponse{
		CategoryName: updateCategory.CategoryName,
		UpdatedBy:    updateCategory.UpdatedBy,
		UserID:       updateCategory.UserID,
		Status:       updateCategory.Status,
		Description:  updateCategory.Description,
		UpdatedAt:    waktuSekarang,
	}
	fmt.Println("Respons:", response)

	return response, nil
}

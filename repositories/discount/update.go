package discount

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *discountRepository) Update(req dto.DiscountUpdateRequest) (dto.DiscountUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateDiscount := entity.Discount{
		DiscountName:       req.DiscountName,
		DiscountPercentage: req.DiscountPercentage,
		Status:             req.Status, // Pastikan status boolean diterima dengan benar
		UpdatedBy:          req.UpdatedBy,
		UserID:             req.UserID,
		Description:        req.Description,
		MerchantID:         req.MerchantID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingDiscount entity.Discount
	err := b.DB.Where("id = ?", req.ID).First(&existingDiscount).Error
	if err != nil {
		return dto.DiscountUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingDiscount).Updates(updateDiscount).Error
	if err != nil {
		return dto.DiscountUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.DiscountUpdateResponse{
		DiscountName:       updateDiscount.DiscountName,
		DiscountPercentage: updateDiscount.DiscountPercentage,
		Status:             updateDiscount.Status,
		UpdatedBy:          updateDiscount.UpdatedBy,
		UserID:             updateDiscount.UserID,
		MerchantID:         updateDiscount.MerchantID,
		Description:        updateDiscount.Description,
	}

	return response, nil
}

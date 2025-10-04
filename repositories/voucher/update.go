package voucher

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *voucherRepository) Update(req dto.VoucherUpdateRequest) (dto.VoucherUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateMerk := entity.VoucherGenerate{
		Status: req.Status,
	}

	var existingProduct entity.VoucherGenerate
	err := b.DB.Where("id = ? AND merchant_id = ?", req.ID, req.MerchantID).First(&existingProduct).Error
	if err != nil {
		return dto.VoucherUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingProduct).Updates(updateMerk).Error
	if err != nil {
		return dto.VoucherUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.VoucherUpdateResponse{
		Status: updateMerk.Status,
	}

	return response, nil
}

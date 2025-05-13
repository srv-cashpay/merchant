package payment

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *paymentRepository) Update(req dto.PaymentUpdateRequest) (dto.PaymentUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updatePayment := entity.Payment{
		PaymentName:       req.PaymentName,
		PaymentPercentage: req.PaymentPercentage,
		Status:            req.Status, // Pastikan status boolean diterima dengan benar
		UpdatedBy:         req.UpdatedBy,
		UserID:            req.UserID,
		Description:       req.Description,
		MerchantID:        req.MerchantID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingPayment entity.Payment
	err := b.DB.Where("id = ?", req.ID).First(&existingPayment).Error
	if err != nil {
		return dto.PaymentUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingPayment).Updates(updatePayment).Error
	if err != nil {
		return dto.PaymentUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.PaymentUpdateResponse{
		PaymentName:       updatePayment.PaymentName,
		PaymentPercentage: updatePayment.PaymentPercentage,
		Status:            updatePayment.Status,
		UpdatedBy:         updatePayment.UpdatedBy,
		UserID:            updatePayment.UserID,
		MerchantID:        updatePayment.MerchantID,
		Description:       updatePayment.Description,
	}

	return response, nil
}

package reservation

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *reservationRepository) Update(req dto.UpdateReservationRequest) (dto.UpdateReservationResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateReservation := entity.Reservation{
		Table:      req.Table,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingReservation entity.Reservation
	err := b.DB.Where("id = ?", req.ID).First(&existingReservation).Error
	if err != nil {
		return dto.UpdateReservationResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingReservation).Updates(updateReservation).Error
	if err != nil {
		return dto.UpdateReservationResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.UpdateReservationResponse{
		Table:      updateReservation.Table,
		UserID:     updateReservation.UserID,
		MerchantID: updateReservation.MerchantID,
	}

	return response, nil
}

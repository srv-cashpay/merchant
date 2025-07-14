package reservation

import (
	"encoding/json"
	"fmt"

	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *reservationRepository) Update(req dto.UpdateReservationRequest) (dto.UpdateReservationResponse, error) {
	tableJSON, err := json.Marshal(req.Table)
	if err != nil {
		return dto.UpdateReservationResponse{}, fmt.Errorf("gagal mengonversi table ke JSON: %w", err)
	}

	updateReservation := entity.Reservation{
		Table:      tableJSON,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		Name:       req.Name,
		Whatsapp:   req.Whatsapp,
		Date:       req.Date,
		Time:       req.Time,
		CreatedBy:  req.CreatedBy,
		DeletedBy:  req.DeletedBy,
	}

	// Cek apakah reservasi ada terlebih dahulu
	var existingReservation entity.Reservation
	err = b.DB.Where("id = ?", req.ID).First(&existingReservation).Error
	if err != nil {
		return dto.UpdateReservationResponse{}, err
	}

	// Update reservasi
	err = b.DB.Model(&existingReservation).Updates(updateReservation).Error
	if err != nil {
		return dto.UpdateReservationResponse{}, err
	}

	// Response setelah update
	response := dto.UpdateReservationResponse{
		ID:         existingReservation.ID,
		UserID:     updateReservation.UserID,
		MerchantID: updateReservation.MerchantID,
		Name:       updateReservation.Name,
		Whatsapp:   updateReservation.Whatsapp,
		Date:       updateReservation.Date,
		Time:       updateReservation.Time,
		Table:      req.Table, // ✅ Gunakan data request yang sudah dalam bentuk []TableRequest
		CreatedBy:  updateReservation.CreatedBy,
		DeletedBy:  updateReservation.DeletedBy,
		CreatedAt:  existingReservation.CreatedAt,
	}

	return response, nil
}

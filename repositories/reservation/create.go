package reservation

import (
	"crypto/rand"
	"encoding/json"
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *reservationRepository) Create(req dto.ReservationRequest) (dto.ReservationResponse, error) {
	// Insert or update the auto_increment value based on merchant_id
	var autoIncrement int
	err := r.DB.Raw(`
		INSERT INTO merchant_auto_increments (merchant_id, next_increment)
		VALUES (?, 1)
		ON CONFLICT (merchant_id) DO UPDATE
		SET next_increment = merchant_auto_increments.next_increment + 1
		RETURNING next_increment - 1;
	`, req.MerchantID).Scan(&autoIncrement).Error

	if err != nil {
		return dto.ReservationResponse{}, err
	}
	productsJSON, err := json.Marshal(req.Table)
	if err != nil {
		return dto.ReservationResponse{}, fmt.Errorf("gagal mengonversi produk ke JSON: %w", err)
	}
	// Create the new reservation entry
	create := entity.Reservation{
		ID:         req.ID,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		Table:      productsJSON,
		CreatedBy:  req.CreatedBy,
		Name:       req.Name,
		Whatsapp:   req.Whatsapp,
		Date:       req.Date,
		Time:       req.Time,
	}

	// Save the new reservation to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.ReservationResponse{}, err
	}
	var responseReservation []dto.TableResponse
	if err := json.Unmarshal(productsJSON, &responseReservation); err != nil {
		return dto.ReservationResponse{}, fmt.Errorf("gagal mengurai JSON produk untuk response: %w", err)
	}
	// Build the response for the created reservation
	response := dto.ReservationResponse{
		ID:         create.ID,
		UserID:     create.UserID,
		MerchantID: create.MerchantID,
		CreatedBy:  create.CreatedBy,
		Name:       create.Name,
		Whatsapp:   create.Whatsapp,
		Date:       create.Date,
		Time:       create.Time,
		Table:      responseReservation,
	}

	return response, nil
}

// Function to generate the reservation ID
func generateReservationID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

	// Generate a secure random part of the reservation ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final reservation ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the reservation ID
func generateSecurePart() (string, error) {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

	// Generate a random string of length 12
	securePart := make([]byte, 12)
	_, err := rand.Read(securePart)
	if err != nil {
		return "", err
	}

	// Map each byte to a character from the chars string
	for i := range securePart {
		securePart[i] = chars[securePart[i]%byte(len(chars))]
	}

	return string(securePart), nil
}

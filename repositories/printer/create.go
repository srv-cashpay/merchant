package printer

import (
	"crypto/rand"
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *printerRepository) Create(req dto.PrinterRequest) (dto.PrinterResponse, error) {
	// Insert or update the auto_increment value based on merchant_id
	var autoIncrement int
	err := r.DB.Raw(`
		INSERT INTO printer_auto_increments (merchant_id, next_increment)
		VALUES (?, 2)
		ON CONFLICT (merchant_id) DO UPDATE
		SET next_increment = printer_auto_increments.next_increment + 1
		RETURNING next_increment - 1;
	`, req.MerchantID).Scan(&autoIncrement).Error

	if err != nil {
		return dto.PrinterResponse{}, err
	}

	// Generate Printer ID with prefix and auto increment value
	prefix := "p="
	secureID, err := generatePrinterID(prefix, autoIncrement)
	if err != nil {
		return dto.PrinterResponse{}, err
	}

	// Create the new printer entry
	create := entity.Printer{
		ID:          secureID,
		PrinterName: req.PrinterName,
		UserID:      req.UserID,
		MerchantID:  req.MerchantID,
		CreatedBy:   req.CreatedBy,
	}

	// Save the new printer to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.PrinterResponse{}, err
	}

	// Build the response for the created printer
	response := dto.PrinterResponse{
		ID:          create.ID,
		UserID:      create.UserID,
		PrinterName: create.PrinterName,
		MerchantID:  create.MerchantID,
		CreatedBy:   create.CreatedBy,
	}

	return response, nil
}

// Function to generate the printer ID
func generatePrinterID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

	// Generate a secure random part of the printer ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final printer ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the printer ID
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

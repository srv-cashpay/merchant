package Sidebar

import (
	"crypto/rand"
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *SidebarRepository) Create(req dto.SidebarRequest) (dto.SidebarResponse, error) {

	// Create the new Sidebar entry
	create := entity.Sidebar{
		ID:         req.ID,
		Label:      req.Label,
		Icon:       req.Icon,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		CreatedBy:  req.CreatedBy,
		To:         req.To,
	}

	// Save the new Sidebar to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.SidebarResponse{}, err
	}

	// Build the response for the created Sidebar
	response := dto.SidebarResponse{
		ID:         create.ID,
		UserID:     create.UserID,
		MerchantID: create.MerchantID,
		Label:      create.Label,
		Icon:       create.Icon,
		To:         create.To,
		CreatedBy:  create.CreatedBy,
	}

	return response, nil
}

// Function to generate the Sidebar ID
func generateSidebarID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

	// Generate a secure random part of the Sidebar ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final Sidebar ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the Sidebar ID
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

package user

import (
	"crypto/rand"
	"fmt"

	"github.com/srv-cashpay/auth/entity"
	dto "github.com/srv-cashpay/merchant/dto"
)

func (r *userRepository) Create(req dto.UserMerchantRequest) (dto.UserMerchantResponse, error) {
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
		return dto.UserMerchantResponse{}, err
	}

	// Generate UserMerchant ID with prefix and auto increment value
	prefix := "p="
	secureID, err := generateUserMerchantID(prefix, autoIncrement)
	if err != nil {
		return dto.UserMerchantResponse{}, err
	}

	// Create the new user entry
	create := entity.AccessDoor{
		ID:           secureID,
		AccessRoleID: req.AccessRoleID,
		FullName:     req.FullName,
		Whatsapp:     req.Whatsapp,
		Email:        req.Email,
		Password:     req.Password,
	}

	// Save the new user to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.UserMerchantResponse{}, err
	}

	// Build the response for the created user
	response := dto.UserMerchantResponse{
		ID:           create.ID,
		AccessRoleID: create.AccessRoleID,
		FullName:     create.FullName,
		Whatsapp:     create.Whatsapp,
		Email:        create.Email,
		Password:     create.Password,
	}

	return response, nil
}

// Function to generate the user ID
func generateUserMerchantID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

	// Generate a secure random part of the user ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final user ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the user ID
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

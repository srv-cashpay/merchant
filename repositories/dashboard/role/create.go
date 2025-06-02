package role

import (
	"crypto/rand"
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *RoleRepository) Create(req dto.RoleRequest) (dto.RoleResponse, error) {
	prefix := "r="

	secureID, err := generateRoleID(prefix)
	if err != nil {
		return dto.RoleResponse{}, err
	}
	// Create the new Role entry
	create := entity.Role{
		ID:   secureID,
		Role: req.Role,
	}

	// Save the new Role to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.RoleResponse{}, err
	}

	// Build the response for the created Role
	response := dto.RoleResponse{
		ID:   secureID,
		Role: req.Role,
	}

	return response, nil
}

func generateRoleID(prefix string) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d")

	// Generate a secure random part of the discount ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final discount ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the discount ID
func generateSecurePart() (string, error) {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

	// Generate a random string of length 12
	securePart := make([]byte, 10)
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

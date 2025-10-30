package user

import (
	"crypto/rand"
	"fmt"
	"time"

	dto "github.com/srv-cashpay/merchant/dto"
	"golang.org/x/crypto/blake2b"
)

func (s *userService) Create(req dto.UserMerchantRequest) (dto.UserMerchantResponse, error) {
	if req.Status != 1 && req.Status != 2 {
		return dto.UserMerchantResponse{}, fmt.Errorf("invalid status: must be 1 (active) or 2 (inactive)")
	}

	create := dto.UserMerchantRequest{
		AccessRoleID: req.AccessRoleID,
		FullName:     req.FullName,
		Whatsapp:     req.Whatsapp,
		Email:        req.Email,
		Password:     req.Password,
		Description:  req.Description,
		Status:       req.Status,
		UserID:       req.UserID,
		MerchantID:   req.MerchantID,
		CreatedBy:    req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.UserMerchantResponse{}, err
	}

	response := dto.UserMerchantResponse{
		AccessRoleID: created.AccessRoleID,
		FullName:     created.FullName,
		Whatsapp:     created.Whatsapp,
		Email:        created.Email,
		Password:     created.Password,
		Description:  created.Description,
		UserID:       created.UserID,
		MerchantID:   created.MerchantID,
		CreatedBy:    created.CreatedBy,
	}

	return response, nil
}

func GenerateSecureID() (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-="

	// Generate a salt
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Combine salt and current timestamp for uniqueness
	timestamp := time.Now().UnixNano()
	saltedID := fmt.Sprintf("%x%d", salt, timestamp)

	// Hash the combination using Blake2
	hash, err := blake2b.New512(nil)
	if err != nil {
		return "", err
	}
	hash.Write([]byte(saltedID))
	hashBytes := hash.Sum(nil)

	// Convert hash bytes into a valid string
	var secureID []byte
	for i := 0; i < 12; i++ {
		secureID = append(secureID, chars[hashBytes[i]%byte(len(chars))])
	}

	return string(secureID), nil
}

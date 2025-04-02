package printer

import (
	"crypto/rand"
	"fmt"
	"time"

	dto "github.com/srv-cashpay/merchant/dto"
	"golang.org/x/crypto/blake2b"
)

func (s *printerService) Create(req dto.PrinterRequest) (dto.PrinterResponse, error) {
	create := dto.PrinterRequest{
		PrinterName: req.PrinterName,
		UserID:      req.UserID,
		ID:          req.ID,
		MerchantID:  req.MerchantID,
		CreatedBy:   req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.PrinterResponse{}, err
	}

	response := dto.PrinterResponse{
		ID:          created.ID,
		UserID:      created.UserID,
		PrinterName: created.PrinterName,
		CreatedBy:   created.CreatedBy,
		MerchantID:  created.MerchantID,
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

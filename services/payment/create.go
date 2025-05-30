package payment

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	"golang.org/x/crypto/blake2b"
	"gorm.io/gorm"
)

func (s *paymentService) Create(req dto.PaymentRequest) (dto.PaymentResponse, error) {
	var merchantDetail entity.MerchantDetail
	err := s.Repo.CheckMerchantDetail(req.MerchantID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.PaymentResponse{}, fmt.Errorf("merchant detail not found for merchant_id: %s", req.MerchantID)
		}
		return dto.PaymentResponse{}, err
	}
	if req.Status != 1 && req.Status != 2 {
		return dto.PaymentResponse{}, fmt.Errorf("invalid status: must be 1 (active) or 2 (inactive)")
	}

	create := dto.PaymentRequest{
		PaymentName: req.PaymentName,
		Status:      req.Status,
		UserID:      req.UserID,
		MerchantID:  req.MerchantID,
		CreatedBy:   req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.PaymentResponse{}, err
	}

	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	// Dapatkan string status berdasarkan nilai integer
	statusString, ok := statusMap[create.Status]
	if !ok {
		return dto.PaymentResponse{}, fmt.Errorf("invalid status value in database")
	}

	response := dto.PaymentResponse{
		ID:          created.ID,
		UserID:      created.UserID,
		PaymentName: created.PaymentName,
		Status:      statusString,
		MerchantID:  created.MerchantID,
		CreatedBy:   created.CreatedBy,
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

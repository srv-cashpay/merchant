package reservation

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"time"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	"golang.org/x/crypto/blake2b"
)

func (s *reservationService) Create(req dto.ReservationRequest) (dto.ReservationResponse, error) {
	tableJSON, err := json.Marshal(req.Table)
	if err != nil {
		return dto.ReservationResponse{}, fmt.Errorf("marshal table gagal: %w", err)
	}

	create := entity.Reservation{
		Name:        req.Name,
		Whatsapp:    req.Whatsapp,
		Date:        req.Date,
		Time:        req.Time,
		Description: req.Description,
		UserID:      req.UserID,
		Table:       tableJSON,
		MerchantID:  req.MerchantID,
		CreatedBy:   req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.ReservationResponse{}, err
	}
	var responseReservation []dto.Table
	if err := json.Unmarshal(tableJSON, &responseReservation); err != nil {
		return dto.ReservationResponse{}, fmt.Errorf("gagal mengurai JSON produk untuk response: %w", err)
	}
	response := dto.ReservationResponse{
		ID:          created.ID,
		UserID:      created.UserID,
		Name:        created.Name,
		Whatsapp:    created.Whatsapp,
		Date:        created.Date,
		Time:        created.Time,
		Description: created.Description,
		MerchantID:  created.MerchantID,
		CreatedBy:   created.CreatedBy,
		Table:       responseReservation,
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

package paymentmethod

import (
	"crypto/rand"
	"fmt"
	"strconv"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	res "github.com/srv-cashpay/util/s/response"
)

func (r *paymentmethodRepository) Create(req dto.PaymentRequest) (dto.PaymentResponse, error) {
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
		return dto.PaymentResponse{}, err
	}

	// Generate Payment ID with prefix and auto increment value
	prefix := "p="
	secureID, err := generatePaymentID(prefix, autoIncrement)
	if err != nil {
		return dto.PaymentResponse{}, err
	}

	// Create the new paymentmethod entry
	create := entity.Payment{
		ID:          secureID,
		PaymentName: req.PaymentName,
		Status:      req.Status,
		UserID:      req.UserID,
		MerchantID:  req.MerchantID,
		CreatedBy:   req.CreatedBy,
	}

	// Save the new paymentmethod to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.PaymentResponse{}, err
	}

	// Map the status from integer to string
	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	createdStatus, err := strconv.Atoi(fmt.Sprintf("%v", create.Status))
	if err != nil {
		return dto.PaymentResponse{}, fmt.Errorf("invalid status value: %v", create.Status)
	}

	statusString, ok := statusMap[createdStatus]
	if !ok {
		return dto.PaymentResponse{}, fmt.Errorf("invalid status value in database")
	}

	// Build the response for the created paymentmethod
	response := dto.PaymentResponse{
		ID:         create.ID,
		UserID:     create.UserID,
		MerchantID: create.MerchantID,
		Status:     statusString,
		CreatedBy:  create.CreatedBy,
	}

	return response, nil
}

// Function to generate the paymentmethod ID
func generatePaymentID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

	// Generate a secure random part of the paymentmethod ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final paymentmethod ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the paymentmethod ID
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

func (r *paymentmethodRepository) CheckMerchantDetail(merchantID string, merchantDetail *entity.MerchantDetail) error {
	err := r.DB.Where("id = ?", merchantID).First(merchantDetail).Error
	if err != nil {
		return err
	}

	// Periksa apakah semua kolom penting sudah terisi
	if merchantDetail.MerchantName == "" || merchantDetail.Address == "" ||
		merchantDetail.Country == "" || merchantDetail.City == "" ||
		merchantDetail.Zip == "" || merchantDetail.Phone == "" {
		return res.ErrorBuilder(&res.ErrorConstant.MerchantNoProvide, nil)
	}

	return nil
}

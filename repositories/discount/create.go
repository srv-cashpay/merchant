package discount

import (
	"crypto/rand"
	"fmt"
	"strconv"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	res "github.com/srv-cashpay/util/s/response"
)

func (r *discountRepository) Create(req dto.DiscountRequest) (dto.DiscountResponse, error) {
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
		return dto.DiscountResponse{}, err
	}

	// Generate Discount ID with prefix and auto increment value
	prefix := "d="
	secureID, err := generateDiscountID(prefix, autoIncrement)
	if err != nil {
		return dto.DiscountResponse{}, err
	}

	// Create the new discount entry
	create := entity.Discount{
		ID:                 secureID,
		DiscountName:       req.DiscountName,
		DiscountPercentage: req.DiscountPercentage,
		Status:             req.Status,
		UserID:             req.UserID,
		MerchantID:         req.MerchantID,
		CreatedBy:          req.CreatedBy,
		Description:        req.Description,
	}

	// Save the new discount to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.DiscountResponse{}, err
	}

	// Map the status from integer to string
	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	createdStatus, err := strconv.Atoi(fmt.Sprintf("%v", create.Status))
	if err != nil {
		return dto.DiscountResponse{}, fmt.Errorf("invalid status value: %v", create.Status)
	}

	statusString, ok := statusMap[createdStatus]
	if !ok {
		return dto.DiscountResponse{}, fmt.Errorf("invalid status value in database")
	}

	// Build the response for the created discount
	response := dto.DiscountResponse{
		ID:                 create.ID,
		UserID:             create.UserID,
		DiscountName:       create.DiscountName,
		DiscountPercentage: create.DiscountPercentage,
		MerchantID:         create.MerchantID,
		Description:        create.Description,
		Status:             statusString,
		CreatedBy:          create.CreatedBy,
	}

	return response, nil
}

// Function to generate the discount ID
func generateDiscountID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

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

func (r *discountRepository) CheckMerchantDetail(merchantID string, merchantDetail *entity.MerchantDetail) error {
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

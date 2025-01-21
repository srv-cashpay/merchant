package category

import (
	"crypto/rand"
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	res "github.com/srv-cashpay/util/s/response"
)

func (r *categoryRepository) Create(req dto.CategoryRequest) (dto.CategoryResponse, error) {
	// Insert or update the auto_increment value based on merchant_id
	var autoIncrement int
	err := r.DB.Raw(`
	INSERT INTO category_auto_increments (category_id, next_increment)
	VALUES (?, 1)
	ON CONFLICT (category_id) DO UPDATE
	SET next_increment = category_auto_increments.next_increment + 1
	RETURNING next_increment - 1;
`, req.MerchantID).Scan(&autoIncrement).Error

	if err != nil {
		return dto.CategoryResponse{}, err
	}

	// Generate Product ID with prefix and auto increment value
	prefix := "c="
	secureID, err := generateProductID(prefix, autoIncrement)
	if err != nil {
		return dto.CategoryResponse{}, err
	}
	create := entity.Category{
		ID:           secureID,
		CategoryName: req.CategoryName,
		MerchantID:   req.MerchantID,
		UserID:       req.UserID,
		CreatedBy:    req.CreatedBy,
		Status:       req.Status,
		Description:  req.Description,
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.CategoryResponse{}, err
	}

	response := dto.CategoryResponse{
		ID:           req.ID,
		CategoryName: create.CategoryName,
		MerchantID:   create.MerchantID,
		UserID:       create.UserID,
		CreatedBy:    create.CreatedBy,
		Status:       create.Status,
		Description:  create.Description,
	}

	return response, nil

}

// Function to generate the product ID
func generateProductID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

	// Generate a secure random part of the product ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final product ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the product ID
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

func (r *categoryRepository) CheckMerchantDetail(merchantID string, merchantDetail *entity.MerchantDetail) error {
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

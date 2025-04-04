package product

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	"gorm.io/gorm"
)

func (s *productService) Upload(req dto.ProductUploadRequest) (dto.ProductUploadResponse, error) {

	// Validasi MerchantDetail
	var merchantDetail entity.MerchantDetail
	err := s.Repo.CheckMerchantDetail(req.MerchantID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.ProductUploadResponse{}, fmt.Errorf("merchant detail not found for merchant_id: %s", req.MerchantID)
		}
		return dto.ProductUploadResponse{}, err
	}

	// Validate file type
	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	ext := strings.ToLower(filepath.Ext(req.File.Filename))
	isAllowed := false
	for _, allowed := range allowedExtensions {
		if ext == allowed {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return dto.ProductUploadResponse{}, errors.New("invalid file type: only JPG and PNG are allowed")
	}

	// Generate unique file name and destination
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	destinationDir := "uploads"
	if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
		return dto.ProductUploadResponse{}, fmt.Errorf("failed to create uploads directory: %w", err)
	}
	fullPath := filepath.Join(destinationDir, newFileName)

	// Set destination in the request
	req.Destination = fullPath

	// Save file and metadata
	response, err := s.Repo.SaveFile(req)
	if err != nil {
		return dto.ProductUploadResponse{}, err
	}

	return response, nil
}

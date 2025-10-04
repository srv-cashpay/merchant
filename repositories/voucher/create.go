package voucher

import (
	"crypto/rand"
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	res "github.com/srv-cashpay/util/s/response"
)

func (r *voucherRepository) Create(req dto.VoucherRequest) (dto.VoucherResponse, error) {
	// Step 1: ambil / update auto increment untuk merchant
	var autoIncrement int
	err := r.DB.Raw(`
		INSERT INTO merchant_auto_increments (merchant_id, next_increment)
		VALUES (?, 1)
		ON CONFLICT (merchant_id) DO UPDATE
		SET next_increment = merchant_auto_increments.next_increment + 1
		RETURNING next_increment - 1;
	`, req.MerchantID).Scan(&autoIncrement).Error
	if err != nil {
		return dto.VoucherResponse{}, err
	}

	// Step 2: simpan voucher (master) -> semua field tersimpan di tabel voucher
	create := entity.Voucher{
		ID:         req.ID,
		MerchantID: req.MerchantID,
		UserID:     req.UserID,
		Nomor:      req.Nomor,
		CreatedBy:  req.CreatedBy,
	}
	if err := r.DB.Create(&create).Error; err != nil {
		return dto.VoucherResponse{}, err
	}

	var dtoGenerates []entity.VoucherGenerate
	for _, v := range req.VoucherGenerate {
		dtoGenerates = append(dtoGenerates, entity.VoucherGenerate{
			ID:          v.ID,
			VoucherID:   create.ID, // FK ke master voucher
			MerchantID:  req.MerchantID,
			VoucherName: v.VoucherName,
			VoucherLink: v.VoucherLink,
			StartDate:   v.StartDate,
			EndDate:     v.EndDate,
			Status:      false,
		})
	}

	if len(dtoGenerates) > 0 {
		if err := r.DB.Create(&dtoGenerates).Error; err != nil {
			return dto.VoucherResponse{}, err
		}
	}

	// Step 4: mapping ke DTO response
	response := dto.VoucherResponse{
		ID:              create.ID,
		MerchantID:      create.MerchantID,
		UserID:          create.UserID,
		CreatedBy:       create.CreatedBy,
		VoucherGenerate: req.VoucherGenerate, // hanya tampil di response
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

func (r *voucherRepository) CheckMerchantDetail(merchantID string, merchantDetail *entity.MerchantDetail) error {
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

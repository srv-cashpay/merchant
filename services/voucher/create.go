package voucher

import (
	"errors"
	"fmt"

	util "github.com/srv-cashpay/util/s"

	"github.com/google/uuid"
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	"gorm.io/gorm"
)

func (s *voucherService) Create(req dto.VoucherRequest) (dto.VoucherResponse, error) {
	// Validasi MerchantDetail
	var merchantDetail entity.MerchantDetail
	err := s.Repo.CheckMerchantDetail(req.MerchantID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.VoucherResponse{}, fmt.Errorf("merchant detail not found for merchant_id: %s", req.MerchantID)
		}
		return dto.VoucherResponse{}, err
	}

	// 🔑 Generate voucher sesuai jumlah "Nomor"
	var voucherGenerates []dto.VoucherGenerate
	for i := 0; i < req.Nomor; i++ {
		secureID, err := util.GenerateProductID()
		if err != nil {
			return dto.VoucherResponse{}, err
		}

		link := "https://cashpay.my.id/voucher-verification/" + secureID + "/" + req.MerchantID

		qr, err := util.GenerateQRCode(link)
		if err != nil {
			return dto.VoucherResponse{}, err
		}

		voucherGenerates = append(voucherGenerates, dto.VoucherGenerate{
			ID:          secureID,
			MerchantID:  req.MerchantID,
			VoucherName: req.VoucherName,
			VoucherLink: link,
			VoucherQR:   qr, // 👈 tambahin field baru di DTO
			StartDate:   req.StartDate,
			EndDate:     req.EndDate,
			Status:      false,
		})
	}

	// Build request untuk Repo.Create
	create := dto.VoucherRequest{
		ID:              uuid.NewString(), // master voucher id
		Nomor:           req.Nomor,
		UserID:          req.UserID,
		MerchantID:      req.MerchantID,
		CreatedBy:       req.CreatedBy,
		VoucherGenerate: voucherGenerates,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.VoucherResponse{}, err
	}

	return created, nil
}

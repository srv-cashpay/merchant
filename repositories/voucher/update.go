package voucher

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *voucherRepository) Update(req dto.VoucherUpdateRequest) (dto.VoucherUpdateResponse, error) {
	// Step 1: Ambil voucher utama
	var existingVoucher entity.Voucher
	err := b.DB.Where("id = ?", req.ID).First(&existingVoucher).Error
	if err != nil {
		return dto.VoucherUpdateResponse{}, err
	}

	// Step 2: Update master voucher
	err = b.DB.Model(&existingVoucher).Updates(entity.Voucher{
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		UpdatedBy:  req.UpdatedBy,
	}).Error
	if err != nil {
		return dto.VoucherUpdateResponse{}, err
	}

	// Step 3: Hapus voucher_generate lama
	err = b.DB.Where("id = ?", req.ID).Delete(&entity.VoucherGenerate{}).Error
	if err != nil {
		return dto.VoucherUpdateResponse{}, err
	}

	// Step 4: Insert voucher_generate baru dari request
	var newGenerates []dto.VoucherGenerate
	for _, vg := range req.VoucherGenerate {
		newEntity := entity.VoucherGenerate{
			MerchantID:  req.MerchantID,
			VoucherName: vg.VoucherName,
			VoucherLink: vg.VoucherLink,
			StartDate:   vg.StartDate,
			EndDate:     vg.EndDate,
			Status:      vg.Status,
		}

		if err := b.DB.Create(&newEntity).Error; err != nil {
			return dto.VoucherUpdateResponse{}, err
		}

		newGenerates = append(newGenerates, dto.VoucherGenerate{
			MerchantID:  newEntity.MerchantID,
			VoucherName: newEntity.VoucherName,
			VoucherLink: newEntity.VoucherLink,
			StartDate:   newEntity.StartDate,
			EndDate:     newEntity.EndDate,
			Status:      newEntity.Status,
		})
	}

	// Step 5: Return response
	response := dto.VoucherUpdateResponse{
		ID:              existingVoucher.ID,
		UserID:          existingVoucher.UserID,
		MerchantID:      existingVoucher.MerchantID,
		UpdatedBy:       existingVoucher.UpdatedBy,
		VoucherGenerate: newGenerates,
	}

	return response, nil
}

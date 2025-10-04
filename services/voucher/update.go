package voucher

import (
	"errors"
	"fmt"

	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	"gorm.io/gorm"
)

func (b *voucherService) Update(req dto.VoucherUpdateRequest) (dto.VoucherUpdateResponse, error) {
	// Step 1: validasi merchant
	var merchantDetail entity.MerchantDetail
	err := b.Repo.CheckMerchantDetail(req.MerchantID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.VoucherUpdateResponse{}, fmt.Errorf("merchant detail not found for merchant_id: %s", req.MerchantID)
		}
		return dto.VoucherUpdateResponse{}, err
	}

	// Step 2: panggil repo update
	updated, err := b.Repo.Update(req)
	if err != nil {
		return dto.VoucherUpdateResponse{}, err
	}

	// Step 3: mapping hasil ke response
	response := dto.VoucherUpdateResponse{
		Status: updated.Status,
	}

	return response, nil
}

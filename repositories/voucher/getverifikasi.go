package voucher

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *voucherRepository) GetVerifikasi(req dto.GetVerifikasi) (*dto.VoucherResponse, error) {
	var tr entity.Voucher
	if err := b.DB.Where("id = ?", req.ID).First(&tr).Error; err != nil {
		return nil, err
	}

	// ambil voucher_generate (entity)
	var generates []entity.VoucherGenerate
	if err := b.DB.Where("voucher_id = ?", tr.ID).Find(&generates).Error; err != nil {
		return nil, err
	}

	// mapping ke DTO
	dtoGenerates := make([]dto.VoucherGenerate, 0, len(generates))
	for _, g := range generates {
		dtoGenerates = append(dtoGenerates, dto.VoucherGenerate{
			MerchantID:  g.MerchantID,
			VoucherName: g.VoucherName,
			VoucherLink: g.VoucherLink,
			StartDate:   g.StartDate,
			EndDate:     g.EndDate,
			Status:      g.Status,
		})
	}

	// mapping ke response
	response := &dto.VoucherResponse{
		ID:              tr.ID,
		UserID:          tr.UserID,
		MerchantID:      tr.MerchantID,
		CreatedBy:       tr.CreatedBy,
		VoucherGenerate: dtoGenerates,
	}

	return response, nil
}

package voucher

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *voucherRepository) GetVerifikasi(req dto.GetVerifikasi) (*dto.GetVerifikasiResponse, error) {
	// ambil voucher_generate
	var generates []entity.VoucherGenerate
	if err := b.DB.Where("id = ? AND merchant_id = ?", req.ID, req.MerchantID).Find(&generates).Error; err != nil {
		return nil, err
	}

	// mapping ke DTO
	dtoGenerates := make([]dto.VoucherGenerate, 0, len(generates))
	for _, g := range generates {
		dtoGenerates = append(dtoGenerates, dto.VoucherGenerate{
			ID:          g.ID,
			MerchantID:  g.MerchantID,
			VoucherName: g.VoucherName,
			VoucherLink: g.VoucherLink,
			StartDate:   g.StartDate,
			EndDate:     g.EndDate,
			Status:      g.Status,
		})
	}

	response := &dto.GetVerifikasiResponse{
		VoucherGenerate: dtoGenerates,
	}

	return response, nil
}

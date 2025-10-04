package voucher

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *voucherRepository) GetById(req dto.GetByIdRequest) (*dto.VoucherResponse, error) {
	tr := entity.Voucher{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	var generates []entity.VoucherGenerate
	if err := b.DB.Where("voucher_id = ?", tr.ID).Find(&generates).Error; err != nil {
		return nil, err
	}

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

	response := &dto.VoucherResponse{
		ID:              tr.ID,
		UserID:          tr.UserID,
		MerchantID:      tr.MerchantID,
		CreatedBy:       tr.CreatedBy,
		VoucherGenerate: dtoGenerates,
	}

	return response, nil
}

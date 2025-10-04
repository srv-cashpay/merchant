package voucher

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *voucherRepository) Delete(req dto.DeleteRequest) (dto.DeleteResponse, error) {
	tr := dto.GetByIdRequest{
		ID: req.ID,
	}

	_, err := b.GetById(tr)
	if err != nil {
		return dto.DeleteResponse{}, err
	}

	// Use GORM BeforeDelete hook to set DeletedBy
	if err := b.DB.Where("id = ?", req.ID).Delete(&entity.Voucher{}).Error; err != nil {
		return dto.DeleteResponse{}, err
	}

	response := dto.DeleteResponse{
		DeletedBy: req.DeletedBy,
	}

	return response, nil
}

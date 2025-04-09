package printer

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *printerRepository) Delete(req dto.DeletePrinterRequest) (dto.DeletePrinterResponse, error) {
	tr := dto.GetByIdRequest{
		ID: req.ID,
	}

	_, err := b.GetById(tr)
	if err != nil {
		return dto.DeletePrinterResponse{}, err
	}

	// Use GORM BeforeDelete hook to set DeletedBy
	if err := b.DB.Where("id = ?", req.ID).Delete(&entity.Printer{}).Error; err != nil {
		return dto.DeletePrinterResponse{}, err
	}

	response := dto.DeletePrinterResponse{
		DeletedBy: req.DeletedBy,
	}

	return response, nil
}

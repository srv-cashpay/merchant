package printer

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *printerRepository) GetById(req dto.GetByIdRequest) (*dto.GetPrinterResponse, error) {
	tr := entity.Printer{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.GetPrinterResponse{
		ID:          tr.ID,
		PrinterName: tr.PrinterName,
		UpdatedBy:   tr.UpdatedBy,
	}

	return response, nil
}

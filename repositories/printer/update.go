package printer

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *printerRepository) Update(req dto.UpdatePrinterRequest) (dto.UpdatePrinterResponse, error) {
	tr := dto.GetByIdRequest{
		ID: req.ID,
	}

	request := entity.Printer{
		ID:          tr.ID,
		PrinterName: req.PrinterName,
		UpdatedBy:   req.UpdatedBy,
	}

	mer, err := b.GetById(tr)
	if err != nil {
		return dto.UpdatePrinterResponse{}, err
	}

	err = b.DB.Where("ID = ?", request.ID).Updates(entity.Printer{
		PrinterName: request.PrinterName,
		UpdatedBy:   request.UpdatedBy,
	}).Error
	if err != nil {
		return dto.UpdatePrinterResponse{}, err
	}

	response := dto.UpdatePrinterResponse{
		PrinterName: request.PrinterName,
		UpdatedBy:   request.UpdatedBy,
		ID:          mer.ID,
	}

	return response, nil
}

package printer

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *printerService) Update(req dto.UpdatePrinterRequest) (dto.UpdatePrinterResponse, error) {
	request := dto.UpdatePrinterRequest{
		ID:          req.ID,
		PrinterName: req.PrinterName,
		UpdatedBy:   req.UpdatedBy,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.UpdatePrinterResponse{
		ID:          request.ID,
		PrinterName: request.PrinterName,
		UpdatedBy:   request.UpdatedBy,
	}

	return response, nil
}

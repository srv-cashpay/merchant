package printer

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *printerRepository) Create(req dto.PrinterRequest) (dto.PrinterResponse, error) {

	// Create the new printer entry
	create := entity.Printer{
		ID:          req.ID,
		PrinterName: req.PrinterName,
		UserID:      req.UserID,
		MerchantID:  req.MerchantID,
		CreatedBy:   req.CreatedBy,
	}

	// Save the new printer to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.PrinterResponse{}, err
	}

	// Build the response for the created printer
	response := dto.PrinterResponse{
		ID:          create.ID,
		UserID:      create.UserID,
		PrinterName: create.PrinterName,
		MerchantID:  create.MerchantID,
		CreatedBy:   create.CreatedBy,
	}

	return response, nil
}

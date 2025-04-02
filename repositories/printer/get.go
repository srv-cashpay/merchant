package printer

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *printerRepository) Get(req dto.GetPrinterRequest) (dto.GetPrinterResponse, error) {

	var data entity.Printer

	if err := r.DB.Where("user_id = ?", req.UserID).Find(&data).Error; err != nil {
		return dto.GetPrinterResponse{}, err
	}

	response := dto.GetPrinterResponse{
		ID:          data.ID,
		UserID:      data.UserID,
		PrinterName: data.PrinterName,
		UpdatedBy:   data.UpdatedBy,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		DeletedAt:   data.DeletedAt,
	}

	return response, nil
}

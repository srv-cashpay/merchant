package paymentmethod

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *paymentmethodRepository) GetPicture(req dto.GetPaymentploadRequest) (*dto.GetPaymentUploadResponse, error) {
	tr := entity.UploadedPayment{
		FileName: req.FileName,
	}

	if err := b.DB.Where("file_name = ?", tr.FileName).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.GetPaymentUploadResponse{
		FilePath: tr.FilePath,
	}

	return response, nil
}

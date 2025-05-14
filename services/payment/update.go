package payment

import "github.com/srv-cashpay/merchant/dto"

func (b *paymentService) Update(req dto.PaymentUpdateRequest) (dto.PaymentUpdateResponse, error) {
	request := dto.PaymentUpdateRequest{
		PaymentName:       req.PaymentName,
		PaymentPercentage: req.PaymentPercentage,
		Status:            req.Status,
		UpdatedBy:         req.UpdatedBy,
		UserID:            req.UserID,
		Description:       req.Description,
	}

	payment, err := b.Repo.Update(req)
	if err != nil {
		return payment, err
	}

	response := dto.PaymentUpdateResponse{
		PaymentName:       request.PaymentName,
		PaymentPercentage: request.PaymentPercentage,
		Status:            request.Status,
		UpdatedBy:         request.UpdatedBy,
		UserID:            request.UserID,
		Description:       request.Description,
	}

	return response, nil
}

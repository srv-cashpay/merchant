package discount

import "github.com/srv-cashpay/merchant/dto"

func (b *discountService) Update(req dto.DiscountUpdateRequest) (dto.DiscountUpdateResponse, error) {
	request := dto.DiscountUpdateRequest{
		DiscountName:       req.DiscountName,
		DiscountPercentage: req.DiscountPercentage,
		Status:             req.Status,
		UpdatedBy:          req.UpdatedBy,
		UserID:             req.UserID,
		Description:        req.Description,
	}

	discount, err := b.Repo.Update(req)
	if err != nil {
		return discount, err
	}

	response := dto.DiscountUpdateResponse{
		DiscountName:       request.DiscountName,
		DiscountPercentage: request.DiscountPercentage,
		Status:             request.Status,
		UpdatedBy:          request.UpdatedBy,
		UserID:             request.UserID,
		Description:        request.Description,
	}

	return response, nil
}

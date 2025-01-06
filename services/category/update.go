package category

import "github.com/srv-cashpay/merchant/dto"

func (b *categoryService) Update(req dto.CategoryUpdateRequest) (dto.CategoryUpdateResponse, error) {
	request := dto.CategoryUpdateRequest{
		CategoryName: req.CategoryName,
		UserID:       req.UserID,
		Status:       req.Status,
		Description:  req.Description,
		UpdatedBy:    req.UpdatedBy,
		MerchantID:   req.MerchantID,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.CategoryUpdateResponse{
		CategoryName: request.CategoryName,
		UserID:       request.UserID,
		Status:       request.Status,
		Description:  request.Description,
		UpdatedBy:    request.UpdatedBy,
		MerchantID:   request.MerchantID,
	}

	return response, nil
}

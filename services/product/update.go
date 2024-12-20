package product

import "github.com/srv-cashpay/merchant/dto"

func (b *productService) Update(req dto.ProductUpdateRequest) (dto.ProductUpdateResponse, error) {
	request := dto.ProductUpdateRequest{
		ProductName:  req.ProductName,
		Stock:        req.Stock,
		MinimalStock: req.MinimalStock,
		Price:        req.Price,
		Status:       req.Status,
		UpdatedBy:    req.UpdatedBy,
		UserID:       req.UserID,
		Description:  req.Description,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.ProductUpdateResponse{
		ProductName:  request.ProductName,
		Stock:        request.Stock,
		MinimalStock: request.MinimalStock,
		Price:        request.Price,
		Status:       request.Status,
		UpdatedBy:    request.UpdatedBy,
		UserID:       request.UserID,
		Description:  request.Description,
	}

	return response, nil
}

package category

import (
	dto "github.com/srv-cashpay/merchant/dto"
	util "github.com/srv-cashpay/util/s"
)

func (s *categoryService) Create(req dto.CategoryRequest) (dto.CategoryResponse, error) {

	create := dto.CategoryRequest{
		ID:           util.GenerateRandomString(),
		UserID:       req.UserID,
		MerchantID:   req.MerchantID,
		CategoryName: req.CategoryName,
		CreatedBy:    req.CreatedBy,
		Status:       req.Status,
		Description:  req.Description,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	response := dto.CategoryResponse{
		ID:           created.ID,
		MerchantID:   created.MerchantID,
		CategoryName: created.CategoryName,
		UserID:       created.UserID,
		CreatedBy:    created.CreatedBy,
		Status:       created.Status,
		Description:  created.Description,
	}

	return response, nil
}

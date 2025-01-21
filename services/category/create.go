package category

import (
	"errors"
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	util "github.com/srv-cashpay/util/s"
	"gorm.io/gorm"
)

func (s *categoryService) Create(req dto.CategoryRequest) (dto.CategoryResponse, error) {
	var merchantDetail entity.MerchantDetail
	err := s.Repo.CheckMerchantDetail(req.MerchantID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.CategoryResponse{}, fmt.Errorf("merchant detail not found for merchant_id: %s", req.MerchantID)
		}
		return dto.CategoryResponse{}, err
	}

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

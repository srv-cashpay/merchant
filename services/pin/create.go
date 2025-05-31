package pin

import (
	"errors"
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	util "github.com/srv-cashpay/util/s"
	"gorm.io/gorm"
)

func (s *pinService) Create(req dto.PinRequest) (dto.PinResponse, error) {
	// Validasi MerchantDetail
	var merchantDetail entity.MerchantDetail
	err := s.Repo.CheckMerchantDetail(req.MerchantID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.PinResponse{}, fmt.Errorf("merchant detail not found for merchant_id: %s", req.MerchantID)
		}
		return dto.PinResponse{}, err
	}

	// Proses pembuatan data Pin
	create := dto.PinRequest{
		ID:         util.GenerateRandomString(),
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		Pin:        req.Pin,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.PinResponse{}, err
	}

	response := dto.PinResponse{
		ID:         created.ID,
		MerchantID: created.MerchantID,
		Pin:        created.Pin,
		UserID:     created.UserID,
		CreatedBy:  created.CreatedBy,
	}

	return response, nil
}

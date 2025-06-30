package order

import (
	"encoding/json"
	"errors"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	res "github.com/srv-cashpay/util/s/response"
	"gorm.io/gorm"
)

func (s *orderService) Order(req dto.OrderRequest) (dto.OrderResponse, error) {

	// Validasi MerchantDetail
	var merchantDetail entity.MerchantDetail
	err := s.Repo.CheckMerchantDetail(req.MerchantID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.OrderResponse{}, (&res.ErrorConstant.NotFound)
		}
		return dto.OrderResponse{}, err

	}
	productJSON, err := json.Marshal(req.Product)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	// Buat copy request dengan string product (opsional kalau mau simpan)
	reqWithJSON := req
	reqWithJSON.ProductJSON = string(productJSON)

	// Kirim ke repo
	created, err := s.Repo.Order(reqWithJSON)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	return created, nil
}

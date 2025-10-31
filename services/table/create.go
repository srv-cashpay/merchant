package table

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *tableService) Create(req dto.TableRequest) (dto.TableResponse, error) {

	create := dto.TableRequest{
		Table:       req.Table,
		Floor:       req.Floor,
		Status:      req.Status,
		Description: req.Description,
		UserID:      req.UserID,
		MerchantID:  req.MerchantID,
		CreatedBy:   req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.TableResponse{}, err
	}

	response := dto.TableResponse{
		ID:          created.ID,
		UserID:      created.UserID,
		Table:       created.Table,
		Floor:       created.Floor,
		Description: created.Description,
		MerchantID:  created.MerchantID,
		CreatedBy:   created.CreatedBy,
	}

	return response, nil
}

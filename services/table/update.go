package table

import "github.com/srv-cashpay/merchant/dto"

func (b *tableService) Update(req dto.TableUpdateRequest) (dto.TableUpdateResponse, error) {
	request := dto.TableUpdateRequest{
		Table:       req.Table,
		Floor:       req.Floor,
		UpdatedBy:   req.UpdatedBy,
		UserID:      req.UserID,
		Description: req.Description,
	}

	table, err := b.Repo.Update(req)
	if err != nil {
		return table, err
	}

	response := dto.TableUpdateResponse{
		Table:       request.Table,
		Floor:       request.Floor,
		UpdatedBy:   request.UpdatedBy,
		UserID:      request.UserID,
		Description: request.Description,
	}

	return response, nil
}

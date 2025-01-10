package table

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *tableRepository) GetById(req dto.GetTableByIdRequest) (*dto.TableResponse, error) {
	tr := entity.Table{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.TableResponse{
		Table: tr.Table,
		Floor: tr.Floor,
	}

	return response, nil
}

package tax

import (
	"fmt"
	"strconv"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *taxRepository) GetById(req dto.GetByIdRequest) (*dto.TaxResponse, error) {
	tr := entity.Tax{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	createdStatus, err := strconv.Atoi(fmt.Sprintf("%v", tr.Status))
	if err != nil {
		return nil, err
	}

	statusString, ok := statusMap[createdStatus]
	if !ok {
		return nil, err
	}

	response := &dto.TaxResponse{
		Tax:           tr.Tax,
		TaxPercentage: tr.TaxPercentage,
		Status:        statusString,
	}

	return response, nil
}

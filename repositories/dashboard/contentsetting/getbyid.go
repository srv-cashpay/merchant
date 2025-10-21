package contentsetting

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *contentsettingRepository) GetById(req dto.GetByContentSettingIdRequest) (*entity.ContentSetting, error) {
	var setting entity.ContentSetting
	if err := r.DB.First(&setting, "id = ?", req.ID).Error; err != nil {
		return nil, err
	}
	return &setting, nil
}

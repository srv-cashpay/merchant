package contentsetting

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *contentsettingRepository) GetById(req dto.GetByContentSettingIdRequest) (*dto.ContentSettingResponse, error) {
	data := entity.ContentSetting{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", data.ID).Take(&data).Error; err != nil {
		return nil, err
	}

	response := &dto.ContentSettingResponse{
		ID:          data.ID,
		UserID:      data.UserID,
		Logo:        data.Logo,
		Title:       data.Title,
		LinkAndroid: data.LinkAndroid,
		LinkApple:   data.LinkApple,
		Description: data.Description,
	}

	return response, nil
}

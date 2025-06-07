package contentsetting

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *contentsettingRepository) Get(req dto.ContentSettingRequest) (dto.ContentSettingResponse, error) {

	var data entity.ContentSetting

	if err := r.DB.Where("user_id = ?", req.UserID).Find(&data).Error; err != nil {
		return dto.ContentSettingResponse{}, err
	}

	response := dto.ContentSettingResponse{
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

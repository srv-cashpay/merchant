package contentsetting

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *contentsettingRepository) Update(req dto.UpdateContentSettingRequest) (dto.UpdateContentSettingResponse, error) {
	data := dto.GetByContentSettingIdRequest{
		ID: req.ID,
	}

	mer, err := b.GetById(data)
	if err != nil {
		return dto.UpdateContentSettingResponse{}, err
	}

	err = b.DB.Where("ID = ?", req.ID).Updates(entity.ContentSetting{
		ID:          data.ID,
		UserID:      mer.UserID,
		Logo:        mer.Logo,
		Title:       mer.Title,
		LinkAndroid: mer.LinkAndroid,
		LinkApple:   mer.LinkApple,
		Description: mer.Description,
	}).Error
	if err != nil {
		return dto.UpdateContentSettingResponse{}, err
	}

	response := dto.UpdateContentSettingResponse{
		ID:          mer.ID,
		UserID:      mer.UserID,
		Logo:        mer.Logo,
		Title:       mer.Title,
		LinkAndroid: mer.LinkAndroid,
		LinkApple:   mer.LinkApple,
		Description: mer.Description,
	}

	return response, nil
}

package contentsetting

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *contentsettingRepository) Update(req dto.UpdateContentSettingRequest) (dto.UpdateContentSettingResponse, error) {
	data := dto.GetByContentSettingIdRequest{
		ID: req.ID,
	}

	request := entity.ContentSetting{
		ID:          data.ID,
		Logo:        req.Logo,
		Title:       req.Title,
		LinkAndroid: req.LinkAndroid,
		LinkApple:   req.LinkApple,
		Description: req.Description,
	}

	mer, err := b.GetById(data)
	if err != nil {
		return dto.UpdateContentSettingResponse{}, err
	}

	err = b.DB.Where("ID = ?", req.ID).Updates(entity.ContentSetting{
		ID:          data.ID,
		UserID:      request.UserID,
		Logo:        request.Logo,
		Title:       request.Title,
		LinkAndroid: request.LinkAndroid,
		LinkApple:   request.LinkApple,
		Description: request.Description,
	}).Error
	if err != nil {
		return dto.UpdateContentSettingResponse{}, err
	}

	response := dto.UpdateContentSettingResponse{
		ID:          mer.ID,
		UserID:      request.UserID,
		Logo:        request.Logo,
		Title:       request.Title,
		LinkAndroid: request.LinkAndroid,
		LinkApple:   request.LinkApple,
		Description: request.Description,
	}

	return response, nil
}

package contentsetting

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (b *contentsettingService) Update(req dto.UpdateContentSettingRequest) (dto.UpdateContentSettingResponse, error) {
	request := dto.UpdateContentSettingRequest{
		ID:          req.ID,
		Logo:        req.Logo,
		Description: req.Description,
		Title:       req.Title,
		LinkAndroid: req.LinkAndroid,
		LinkApple:   req.LinkApple,
		UpdatedBy:   req.UpdatedBy,
	}

	contentsetting, err := b.Repo.Update(req)
	if err != nil {
		return contentsetting, err
	}

	response := dto.UpdateContentSettingResponse{
		ID:          request.ID,
		Logo:        request.Logo,
		Description: request.Description,
		Title:       request.Title,
		LinkAndroid: request.LinkAndroid,
		LinkApple:   request.LinkApple,
		UpdatedBy:   request.UpdatedBy,
	}

	return response, nil
}

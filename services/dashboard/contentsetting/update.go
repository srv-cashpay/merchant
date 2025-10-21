package contentsetting

import (
	"encoding/json"
	"time"

	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *contentsettingService) Update(req dto.UpdateContentSettingRequest) (dto.UpdateContentSettingResponse, error) {
	// Ambil data lama dari DB
	setting, err := s.Repo.GetById(dto.GetByContentSettingIdRequest{ID: req.ID})
	if err != nil {
		return dto.UpdateContentSettingResponse{}, err
	}

	// 🔹 Marshal tiap array struct ke JSON string
	topHeaderJSON, _ := json.Marshal(req.TopHeader)
	buttonHeaderJSON, _ := json.Marshal(req.ButtonHeader)
	featureJSON, _ := json.Marshal(req.Feature)
	footerJSON, _ := json.Marshal(req.Footer)

	// 🔹 Simpan sebagai array string (karena field di entity adalah []string)
	setting.TopHeader = []string{string(topHeaderJSON)}
	setting.ButtonHeader = []string{string(buttonHeaderJSON)}
	setting.Feature = []string{string(featureJSON)}
	setting.Footer = []string{string(footerJSON)}
	setting.UpdatedBy = req.UpdatedBy
	setting.UpdatedAt = time.Now()

	// 🔹 Update ke database
	if err := s.Repo.Update(setting); err != nil {
		return dto.UpdateContentSettingResponse{}, err
	}

	// 🔹 Return response DTO
	return dto.UpdateContentSettingResponse{
		ID:           setting.ID,
		UserID:       setting.UserID,
		MerchantID:   setting.MerchantID,
		TopHeader:    req.TopHeader,
		ButtonHeader: req.ButtonHeader,
		Feature:      req.Feature,
		Footer:       req.Footer,
		UpdatedBy:    setting.UpdatedBy,
		UpdatedAt:    setting.UpdatedAt,
	}, nil
}

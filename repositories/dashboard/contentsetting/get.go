package contentsetting

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *contentsettingRepository) Get(req dto.ContentSettingRequest) (dto.ContentSettingResponse, error) {
	var data entity.ContentSetting

	// ambil data berdasarkan ID (atau bisa pakai merchant_id tergantung kebutuhan)
	if err := r.DB.First(&data, "id = ?", req.ID).Error; err != nil {
		return dto.ContentSettingResponse{}, err
	}

	response := dto.ContentSettingResponse{
		ID:           data.ID,
		UserID:       data.UserID,
		MerchantID:   data.MerchantID,
		TopHeader:    data.TopHeader,
		ButtonHeader: data.ButtonHeader,
		Feature:      data.Feature,
		Footer:       data.Footer,
		UpdatedBy:    data.UpdatedBy,
		UpdatedAt:    data.UpdatedAt,
	}

	return response, nil
}

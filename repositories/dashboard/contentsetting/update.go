package contentsetting

import "github.com/srv-cashpay/merchant/entity"

func (r *contentsettingRepository) Update(setting *entity.ContentSetting) error {
	return r.DB.Save(setting).Error
}

package contentsetting

import (
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req dto.ContentSettingRequest) (dto.ContentSettingResponse, error)
	GetById(req dto.GetByContentSettingIdRequest) (*dto.ContentSettingResponse, error)
	Update(req dto.UpdateContentSettingRequest) (dto.UpdateContentSettingResponse, error)
}

type contentsettingRepository struct {
	DB *gorm.DB
}

func NewContentSettingRepository(DB *gorm.DB) DomainRepository {
	return &contentsettingRepository{
		DB: DB,
	}
}

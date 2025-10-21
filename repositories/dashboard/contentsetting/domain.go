package contentsetting

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req dto.ContentSettingRequest) (dto.ContentSettingResponse, error)
	GetById(req dto.GetByContentSettingIdRequest) (*entity.ContentSetting, error)
	Update(setting *entity.ContentSetting) error
}

type contentsettingRepository struct {
	DB *gorm.DB
}

func NewContentSettingRepository(DB *gorm.DB) DomainRepository {
	return &contentsettingRepository{
		DB: DB,
	}
}

package Sidebar

import (
	dto "github.com/srv-cashpay/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.SidebarRequest) (dto.SidebarResponse, error)
	Get(req dto.SidebarRequest) (dto.GetSidebarResponse, error)
	GetById(req dto.GetSidebarByIdRequest) (*dto.SidebarResponse, error)
	Delete(req dto.DeleteSidebarRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.SidebarUpdateRequest) (dto.SidebarUpdateResponse, error)
}

type SidebarRepository struct {
	DB *gorm.DB
}

func NewSidebarRepository(DB *gorm.DB) DomainRepository {
	return &SidebarRepository{
		DB: DB,
	}
}

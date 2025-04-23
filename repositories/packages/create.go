package packages

import (
	auth "github.com/srv-cashpay/auth/entity"
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (r *packagesRepository) Create(req dto.PackagesRequest) (dto.PackagesResponse, error) {
	// 1. Membuat entitas package berdasarkan request
	create := entity.Package{
		ID:          req.ID,
		UserID:      req.UserID,
		CreatedBy:   req.CreatedBy,
		Status:      req.Status,
		GrossAmount: req.GrossAmount,
		OrderID:     req.OrderID, // Order ID di sini digunakan untuk transaksi Midtrans
	}

	// 2. Simpan package ke database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.PackagesResponse{}, err
	}

	// 3. Return response dengan data yang telah disimpan
	response := dto.PackagesResponse{
		ID:          create.ID,
		UserID:      create.UserID,
		CreatedBy:   create.CreatedBy,
		Status:      create.Status,
		GrossAmount: create.GrossAmount,
		OrderID:     create.OrderID,
	}

	return response, nil
}

func (r *packagesRepository) UpdateStatus(orderID string, status string) error {
	// Update status berdasarkan OrderID
	if err := r.DB.Model(&entity.Package{}).
		Where("order_id = ?", orderID).
		Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *packagesRepository) FindByID(id string) (auth.AccessDoor, error) {
	var user auth.AccessDoor
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return auth.AccessDoor{}, err
	}
	return user, nil
}

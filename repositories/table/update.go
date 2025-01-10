package table

import (
	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *tableRepository) Update(req dto.TableUpdateRequest) (dto.TableUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateTable := entity.Table{
		Table:      req.Table,
		Floor:      req.Floor,
		UpdatedBy:  req.UpdatedBy,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingTable entity.Table
	err := b.DB.Where("id = ?", req.ID).First(&existingTable).Error
	if err != nil {
		return dto.TableUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingTable).Updates(updateTable).Error
	if err != nil {
		return dto.TableUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.TableUpdateResponse{
		Table:      updateTable.Table,
		Floor:      updateTable.Floor,
		UpdatedBy:  updateTable.UpdatedBy,
		UserID:     updateTable.UserID,
		MerchantID: updateTable.MerchantID,
	}

	return response, nil
}

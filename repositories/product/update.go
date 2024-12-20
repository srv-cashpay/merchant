package product

import (
	"log"

	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/product/entity"
)

func (b *productRepository) Update(req dto.ProductUpdateRequest) (dto.ProductUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateProduct := entity.Product{
		ProductName:  req.ProductName,
		Stock:        req.Stock,
		MinimalStock: req.MinimalStock,
		Price:        req.Price,
		Status:       req.Status, // Pastikan status boolean diterima dengan benar
		UpdatedBy:    req.UpdatedBy,
		UserID:       req.UserID,
		Description:  req.Description,
		MerchantID:   req.MerchantID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingProduct entity.Product
	err := b.DB.Where("user_id = ?", req.UserID).First(&existingProduct).Error
	if err != nil {
		return dto.ProductUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingProduct).Updates(updateProduct).Error
	if err != nil {
		return dto.ProductUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.ProductUpdateResponse{
		ProductName:  updateProduct.ProductName,
		Stock:        updateProduct.Stock,
		MinimalStock: updateProduct.MinimalStock,
		Price:        updateProduct.Price,
		Status:       updateProduct.Status,
		UpdatedBy:    updateProduct.UpdatedBy,
		UserID:       updateProduct.UserID,
		MerchantID:   updateProduct.MerchantID,
		Description:  updateProduct.Description,
	}

	// Log status yang diupdate
	log.Printf("Updated status: %v", updateProduct.Status)

	return response, nil
}

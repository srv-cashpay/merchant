package product

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/product/entity"
)

func (b *productRepository) BulkEdit(req dto.BulkEditRequest) (int, error) {
	tx := b.DB.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	var totalUpdated int

	for _, item := range req.Items {
		if item.ID == "" {
			continue
		}

		updates := map[string]interface{}{}

		if item.SKU != nil {
			updates["sku"] = *item.SKU
		}
		if item.Barcode != nil {
			updates["barcode"] = *item.Barcode
		}
		if item.MerkID != nil {
			updates["merk_id"] = *item.MerkID
		}
		if item.CategoryID != nil {
			updates["category_id"] = *item.CategoryID
		}
		if item.ProductName != nil {
			updates["product_name"] = *item.ProductName
		}
		if item.Description != nil {
			updates["description"] = *item.Description
		}
		if item.Stock != nil {
			updates["stock"] = *item.Stock
		}
		if item.MinimalStock != nil {
			updates["minimal_stock"] = *item.MinimalStock
		}
		if item.Price != nil {
			updates["price"] = *item.Price
		}
		if item.Status != nil {
			updates["status"] = *item.Status
		}

		updates["updated_by"] = req.UpdatedBy

		if len(updates) == 0 {
			continue
		}

		result := tx.Model(&entity.Product{}).Where("id = ?", item.ID).Updates(updates)
		if result.Error != nil {
			tx.Rollback()
			return 0, result.Error
		}
		totalUpdated += int(result.RowsAffected)
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return totalUpdated, nil
}

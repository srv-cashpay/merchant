package voucher

import (
	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
)

func (b *voucherRepository) BulkDelete(req dto.BulkDeleteRequest) (int, error) {
	// Hapus semua produk yang sesuai dengan ID
	result := b.DB.Where("id IN ?", req.ID).Delete(&entity.Voucher{})
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil // Mengembalikan jumlah produk yang berhasil dihapus
}

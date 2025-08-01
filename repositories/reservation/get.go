package reservation

import (
	"fmt"
	"math"
	"strings"

	dto "github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/merchant/entity"
	"github.com/srv-cashpay/merchant/helpers"
)

func (r *reservationRepository) Get(req *dto.Pagination) (RepositoryResult, int) {
	var reservations []entity.Reservation

	var totalRows int64
	totalPages, fromRow, toRow := 0, 0, 0

	// Ubah offset agar sesuai dengan page yang dimulai dari 1
	offset := (req.Page - 1) * req.Limit

	// Ambil data sesuai limit, offset, dan urutan
	find := r.DB.Where("merchant_id = ?", req.MerchantID).Limit(req.Limit).Offset(offset).Order(req.Sort)

	// Generate where query untuk search
	if req.Searchs != nil {
		for _, value := range req.Searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				find = find.Where(fmt.Sprintf("%s = ?", column), query)
			case "contains":
				find = find.Where(fmt.Sprintf("%s LIKE ?", column), "%"+query+"%")
			case "in":
				find = find.Where(fmt.Sprintf("%s IN (?)", column), strings.Split(query, ","))
			}
		}
	}

	find = find.Find(&reservations)

	// Periksa jika ada error saat pengambilan data
	if errFind := find.Error; errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	req.Rows = reservations

	// Hitung total data
	if errCount := r.DB.Model(&entity.Reservation{}).Where("merchant_id = ?", req.MerchantID).Count(&totalRows).Error; errCount != nil {
		return RepositoryResult{Error: errCount}, totalPages
	}

	for i := range reservations {
		tableStr := string(reservations[i].Table)
		truncated := helpers.TruncateString(tableStr, 47)
		reservations[i].Table = []byte(truncated)
	}

	req.TotalRows = int(totalRows)

	// Hitung total halaman berdasarkan limit
	totalPages = int(math.Ceil(float64(totalRows) / float64(req.Limit)))
	req.TotalPages = totalPages
	// Hitung `fromRow` dan `toRow` untuk page saat ini
	if req.Page == 1 {
		// Untuk halaman pertama
		fromRow = 1
		toRow = req.Limit
	} else {
		if req.Page <= totalPages {
			fromRow = (req.Page-1)*req.Limit + 1
			toRow = req.Page * req.Limit
		}
	}

	// Pastikan `toRow` tidak melebihi `totalRows`
	if toRow > int(totalRows) {
		toRow = int(totalRows)
	}

	// Set hasil akhir
	req.FromRow = fromRow
	req.ToRow = toRow

	return RepositoryResult{Result: req}, totalPages
}

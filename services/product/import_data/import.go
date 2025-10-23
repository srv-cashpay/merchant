package import_data

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/srv-cashpay/product/entity"
)

func (s *importService) ImportProducts(ctx context.Context, fileHeader *multipart.FileHeader) (map[string]interface{}, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("gagal membuka file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	// Baca header
	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("gagal membaca header: %w", err)
	}

	expectedHeaders := []string{
		"barcode", "sku", "merk_id", "category_id",
		"product_name", "stock", "minimal_stock",
		"price", "description", "status",
	}

	if len(headers) != len(expectedHeaders) {
		return nil, fmt.Errorf("template tidak sesuai: jumlah kolom tidak cocok")
	}

	for i, h := range headers {
		if h != expectedHeaders[i] {
			return nil, fmt.Errorf("kolom ke-%d harus '%s', bukan '%s'", i+1, expectedHeaders[i], h)
		}
	}

	var imported []entity.Product
	rowNumber := 2 // Karena baris pertama adalah header

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("gagal membaca baris %d: %w", rowNumber, err)
		}

		if len(record) < len(expectedHeaders) {
			return nil, fmt.Errorf("baris %d tidak lengkap", rowNumber)
		}

		// Parsing data sesuai kolom
		product := entity.Product{
			ID:          fmt.Sprintf("PROD-%d", time.Now().UnixNano()),
			Barcode:     record[0],
			UserID:      "", // bisa isi dari context jwt
			MerchantID:  "", // bisa isi dari context jwt
			MerkID:      record[2],
			CategoryID:  record[3],
			ProductName: record[4],
			Description: record[8],
			CreatedBy:   "import_system",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// Convert angka
		if stock, err := strconv.Atoi(record[5]); err == nil {
			product.Stock = stock
		}
		if minStock, err := strconv.Atoi(record[6]); err == nil {
			product.MinimalStock = minStock
		}
		if price, err := strconv.Atoi(record[7]); err == nil {
			product.Price = price
		}
		if status, err := strconv.Atoi(record[9]); err == nil {
			product.Status = status
		}

		imported = append(imported, product)
		rowNumber++
	}

	// Simpan ke DB via repository
	if err := s.Repo.BulkInsert(ctx, imported); err != nil {
		return nil, fmt.Errorf("gagal menyimpan data produk: %w", err)
	}

	return map[string]interface{}{
		"message":       fmt.Sprintf("%d produk berhasil diimport", len(imported)),
		"importedCount": len(imported),
	}, nil
}

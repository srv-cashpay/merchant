package import_data

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/srv-cashpay/product/entity"
	"github.com/xuri/excelize/v2"
)

func (s *importService) ImportProducts(ctx context.Context, fileHeader *multipart.FileHeader) (map[string]interface{}, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("gagal membuka file: %v", err)
	}
	defer file.Close()

	// Baca seluruh isi file Excel
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca file: %v", err)
	}

	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("file bukan format Excel yang valid (.xlsx)")
	}

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca sheet: %v", err)
	}

	if len(rows) <= 1 {
		return nil, fmt.Errorf("template kosong atau tidak ada data")
	}

	var imported int
	for i, row := range rows[1:] { // Skip header
		if len(row) < 13 {
			continue
		}

		// Ambil kolom dari Excel (A-M)
		id := row[0]
		merchantID := row[1]
		userID := row[2]
		barcode := row[3]
		sku := parseUint(row[4])
		merkID := row[5]
		categoryID := row[6]
		productName := row[7]
		stock, _ := strconv.Atoi(row[8])
		minStock, _ := strconv.Atoi(row[9])
		price, _ := strconv.Atoi(row[10])
		description := row[11]
		status, _ := strconv.Atoi(row[12])

		// Jika kolom ID kosong → generate ID unik
		if id == "" {
			newID, err := generateProductID("p=")
			if err != nil {
				return nil, fmt.Errorf("gagal membuat secure ID: %v", err)
			}
			id = newID
		}

		product := entity.Product{
			ID:           id,
			MerchantID:   merchantID,
			UserID:       userID,
			Barcode:      barcode,
			SKU:          sku,
			MerkID:       merkID,
			CategoryID:   categoryID,
			ProductName:  productName,
			Stock:        stock,
			MinimalStock: minStock,
			Price:        price,
			Description:  description,
			Status:       status,
		}

		if err := s.Repo.Create(ctx, &product); err != nil {
			return nil, fmt.Errorf("baris %d gagal disimpan: %v", i+2, err)
		}

		imported++
	}

	return map[string]interface{}{
		"message":       "import berhasil",
		"importedCount": imported,
	}, nil
}

// ===========================
// 🔹 Helper Functions
// ===========================

func parseUint(s string) uint64 {
	v, _ := strconv.ParseUint(s, 10, 64)
	return v
}

func generateProductID(prefix string) (string, error) {
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", prefix, securePart), nil
}

func generateSecurePart() (string, error) {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

	securePart := make([]byte, 12)
	_, err := rand.Read(securePart)
	if err != nil {
		return "", err
	}

	for i := range securePart {
		securePart[i] = chars[securePart[i]%byte(len(chars))]
	}

	return string(securePart), nil
}

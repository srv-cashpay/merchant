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

	merchantID := ctx.Value("MerchantId").(string)
	userID := ctx.Value("UserId").(string)
	createdBy := ctx.Value("CreatedBy").(string)

	var imported int
	for i, row := range rows[1:] { // skip header
		if len(row) < 9 { // karena mulai dari Barcode, SKU, dst
			continue
		}

		barcode := row[0]
		sku := parseUint(row[1])
		merkID := row[2]
		categoryID := row[3]
		productName := row[4]
		stock, _ := strconv.Atoi(row[5])
		minStock, _ := strconv.Atoi(row[6])
		price, _ := strconv.Atoi(row[7])
		description := row[8]
		status := 1
		if len(row) > 9 {
			status, _ = strconv.Atoi(row[9])
		}

		newID, err := generateProductID("p=")
		if err != nil {
			return nil, fmt.Errorf("gagal membuat ID produk di baris %d: %v", i+2, err)
		}

		product := entity.Product{
			ID:           newID,
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
			CreatedBy:    createdBy,
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

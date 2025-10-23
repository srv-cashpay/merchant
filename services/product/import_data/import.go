package import_data

import (
	"bytes"
	"context"
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

	// Read Excel
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca file: %v", err)
	}

	f, err := excelize.OpenReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("file bukan format excel yang valid")
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
	for i, row := range rows[1:] { // skip header
		if len(row) < 10 {
			continue
		}

		stock, _ := strconv.Atoi(row[5])
		minStock, _ := strconv.Atoi(row[6])
		price, _ := strconv.Atoi(row[7])
		status, _ := strconv.Atoi(row[9])

		product := entity.Product{
			Barcode:      row[0],
			MerkID:       row[2],
			CategoryID:   row[3],
			ProductName:  row[4],
			Stock:        stock,
			MinimalStock: minStock,
			Price:        price,
			Description:  row[8],
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

func parseUint(s string) uint64 {
	v, _ := strconv.ParseUint(s, 10, 64)
	return v
}

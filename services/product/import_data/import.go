package import_data

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/srv-cashpay/merchant/dto"
	"github.com/srv-cashpay/product/entity"
)

func (s *importService) ImportProducts(ctx context.Context, fileHeader *multipart.FileHeader, userID string) (dto.ImportResultDTO, error) {
	if fileHeader == nil {
		return dto.ImportResultDTO{}, errors.New("file tidak ditemukan")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return dto.ImportResultDTO{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return dto.ImportResultDTO{}, errors.New("gagal membaca CSV: " + err.Error())
	}

	if len(rows) < 2 {
		return dto.ImportResultDTO{}, errors.New("file tidak memiliki data")
	}

	var (
		validProducts []entity.Product
		failedRows    []string
	)

	for i, row := range rows[1:] {
		if len(row) < 7 {
			failedRows = append(failedRows, fmt.Sprintf("baris %d tidak lengkap", i+2))
			continue
		}

		stock, _ := strconv.Atoi(row[3])
		minStock, _ := strconv.Atoi(row[4])
		price, _ := strconv.Atoi(row[5])

		if row[2] == "" {
			failedRows = append(failedRows, fmt.Sprintf("baris %d: nama produk kosong", i+2))
			continue
		}

		product := entity.Product{
			ID:           fmt.Sprintf("PROD-%d", time.Now().UnixNano()),
			Barcode:      row[0],
			CategoryID:   row[1],
			ProductName:  row[2],
			Stock:        stock,
			MinimalStock: minStock,
			Price:        price,
			Description:  row[6],
			UserID:       userID,
			Status:       1,
			CreatedBy:    userID,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		validProducts = append(validProducts, product)
	}

	if len(validProducts) == 0 {
		return dto.ImportResultDTO{
			Total:    len(rows) - 1,
			Success:  0,
			Failed:   len(failedRows),
			Failures: failedRows,
		}, errors.New("tidak ada data valid untuk disimpan")
	}

	err = s.Repo.SaveBatch(ctx, validProducts)
	if err != nil {
		return dto.ImportResultDTO{}, err
	}

	return dto.ImportResultDTO{
		Total:    len(rows) - 1,
		Success:  len(validProducts),
		Failed:   len(failedRows),
		Failures: failedRows,
	}, nil
}

package import_data

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"time"
)

func (s *importService) GenerateTemplate() ([]byte, string, error) {
	headers := []string{
		"barcode", "sku", "merk_id", "category_id",
		"product_name", "stock", "minimal_stock",
		"price", "description", "status",
	}

	example := []string{
		"BR12345", "1001", "MERK01", "CAT01",
		"Nasi Goreng", "100", "10", "15000",
		"Menu populer", "1",
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	if err := writer.Write(headers); err != nil {
		return nil, "", err
	}
	if err := writer.Write(example); err != nil {
		return nil, "", err
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, "", err
	}

	filename := fmt.Sprintf("product_template_%d.csv", time.Now().Unix())
	return buf.Bytes(), filename, nil
}

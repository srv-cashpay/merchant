package import_data

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

func (s *importService) GenerateTemplate() ([]byte, string, error) {
	f := excelize.NewFile()
	sheet := "Template"
	f.NewSheet(sheet)
	f.DeleteSheet("Sheet1")

	// Header columns
	headers := []string{
		"Barcode", "SKU", "Merk ID", "Category ID",
		"Product Name", "Stock", "Minimal Stock",
		"Price", "Description", "Status",
	}

	// ✅ Gunakan struct excelize.Style, bukan string
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true},
		Alignment: &excelize.Alignment{Horizontal: "center"},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#E2EFDA"}, Pattern: 1},
	})

	for i, h := range headers {
		col := string(rune('A' + i))
		cell := fmt.Sprintf("%s1", col)
		f.SetCellValue(sheet, cell, h)
		f.SetCellStyle(sheet, cell, cell, headerStyle)
	}

	// Example row
	example := []interface{}{
		"BR12345", 1001, "MERK01", "CAT01", "Nasi Goreng", 100, 10, 15000, "Menu populer", 1,
	}
	for i, val := range example {
		col := string(rune('A' + i))
		f.SetCellValue(sheet, fmt.Sprintf("%s2", col), val)
	}

	f.SetColWidth(sheet, "A", "J", 18)

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, "", err
	}

	filename := fmt.Sprintf("product_template_%d.xlsx", time.Now().Unix())
	return buf.Bytes(), filename, nil
}

package import_data

import (
	"bytes"

	"github.com/xuri/excelize/v2"
)

func (s *importService) GenerateTemplate() ([]byte, string, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	headers := []string{
		"Barcode", "SKU", "Merk ID", "Category ID", "Product Name",
		"Stock", "Minimal Stock", "Price", "Description", "Status",
	}

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	style, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	f.SetCellStyle(sheet, "A1", "J1", style)

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, "", err
	}

	return buf.Bytes(), "product_template.xlsx", nil
}

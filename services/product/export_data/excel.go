package export_data

import (
	"context"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func (s *exportService) ExportExcel(ctx context.Context) (*excelize.File, error) {
	products, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "Product Export"
	f.SetSheetName("Sheet1", sheet)
	f.SetSheetRow(sheet, "A1", &[]string{"ID", "Product Name", "Price", "Stock"})

	for i, p := range products {
		row := []interface{}{p.ID, p.ProductName, p.Price, p.Stock}
		cell := fmt.Sprintf("A%d", i+2)
		f.SetSheetRow(sheet, cell, &row)
	}

	return f, nil
}

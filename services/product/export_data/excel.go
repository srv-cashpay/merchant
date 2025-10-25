package export_data

import (
	"context"
	"fmt"

	dto "github.com/srv-cashpay/merchant/dto"

	"github.com/xuri/excelize/v2"
)

func (s *exportService) ExportExcel(ctx context.Context, req dto.ExportFilter) (*excelize.File, error) {
	products, err := s.Repo.FindByFilter(ctx, req)
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "Product Export"
	f.SetSheetName("Sheet1", sheet)
	f.SetSheetRow(sheet, "A1", &[]string{"ID", "Product Name", "Price", "Stock", "Created At"})

	for i, p := range products {
		row := []interface{}{p.ID, p.ProductName, p.Price, p.Stock, p.CreatedAt.Format("2006-01-02")}
		cell := fmt.Sprintf("A%d", i+2)
		f.SetSheetRow(sheet, cell, &row)
	}

	return f, nil
}

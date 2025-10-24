package export_data

import (
	"context"

	"github.com/xuri/excelize/v2"
)

func (s *exportService) ExportExcel(ctx context.Context) (*excelize.File, error) {
	users, err := s.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "Users"
	f.NewSheet(sheet)
	f.SetSheetRow(sheet, "A1", &[]string{"ID", "Name", "Price", "Stock"})

	for i, user := range users {
		row := []interface{}{user.ID, user.ProductName, user.Price, user.Stock}
		cell := "A" + string(rune(i+2))
		f.SetSheetRow(sheet, cell, &row)
	}

	return f, nil
}

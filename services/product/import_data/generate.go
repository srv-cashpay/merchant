package import_data

import (
	"bytes"
	"encoding/csv"
)

func (s *importService) GenerateTemplate() ([]byte, string, error) {
	buf := &bytes.Buffer{}
	writer := csv.NewWriter(buf)

	// header
	header := []string{"Name", "Email", "Role"}
	if err := writer.Write(header); err != nil {
		return nil, "", err
	}

	// contoh data
	sample := []string{"John Doe", "john@example.com", "admin"}
	if err := writer.Write(sample); err != nil {
		return nil, "", err
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, "", err
	}

	return buf.Bytes(), "template_users.csv", nil
}

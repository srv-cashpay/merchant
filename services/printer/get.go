package printer

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *printerService) Get(req dto.GetPrinterRequest) (dto.GetPrinterResponse, error) {
	// Fetch comments from the repository layer based on post_id
	comments, err := s.Repo.Get(req)
	if err != nil {
		return dto.GetPrinterResponse{}, err
	}

	return comments, nil
}

package dashboard

import "github.com/srv-cashpay/merchant/dto"

func (s *getMerkdashboardService) Get(req dto.MerkRequest) ([]dto.MerkResponse, error) {
	merk, err := s.Repo.Get(req)
	if err != nil {
		return nil, err
	}

	return merk, nil
}

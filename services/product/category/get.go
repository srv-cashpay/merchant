package dashboard

import "github.com/srv-cashpay/merchant/entity"

func (s *getCategorydashboardService) Get() ([]entity.Category, error) {
	// Fetch comments from the repository layer based on post_id
	comments, err := s.Repo.Get()
	if err != nil {
		return nil, err
	}

	return comments, nil
}

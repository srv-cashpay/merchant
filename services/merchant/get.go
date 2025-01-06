package product

import (
	dto "github.com/srv-cashpay/merchant/dto"
)

func (s *merchantService) Get(req dto.GetMerchantRequest) (dto.GetMerchantResponse, error) {
	// Fetch comments from the repository layer based on post_id
	comments, err := s.Repo.Get(req)
	if err != nil {
		return dto.GetMerchantResponse{}, err
	}

	return comments, nil
}

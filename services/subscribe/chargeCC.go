package subscribe

import dto "github.com/srv-cashpay/merchant/dto"

// TokenizeCard adalah service untuk men-tokenisasi kartu kredit
func (s *subscribeService) TokenizeCard(cardData dto.TokenizeRequest) (*dto.TokenizeResponse, error) {
	// Panggil repository untuk tokenisasi
	response, err := s.Repo.TokenizeCard(cardData)
	if err != nil {
		return nil, err
	}

	// Return transaksi dalam bentuk entity
	return &dto.TokenizeResponse{
		TokenID:       response.TokenID,
		TransactionID: response.TransactionID,
		Status:        response.Status,
	}, nil
}

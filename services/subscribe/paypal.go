package subscribe

import (
	"github.com/srv-cashpay/merchant/dto"
)

func (s *subscribeService) CreatePaypalOrder(req dto.PaypalCreateRequest) (*dto.PaypalOrderResponse, error) {
	order, err := s.Repo.CreatePaypalOrder(req.Amount, req.Currency)
	if err != nil {
		return nil, err
	}

	// Ambil approval URL dari links
	var approvalLink string
	for _, link := range order.Links {
		if link.Rel == "approve" {
			approvalLink = link.Href
			break
		}
	}

	return &dto.PaypalOrderResponse{
		ID:     order.ID,
		Status: order.Status,
		Link:   approvalLink,
	}, nil
}

func (s *subscribeService) CapturePaypalOrder(orderID string) (*dto.PaypalCaptureResponse, error) {
	capture, err := s.Repo.CapturePaypalOrder(orderID)
	if err != nil {
		return nil, err
	}

	return &dto.PaypalCaptureResponse{
		Status: capture.Status,
	}, nil
}

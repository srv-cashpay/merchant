package order

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"firebase.google.com/go/messaging"
	"github.com/srv-cashpay/merchant/dto"
)

func (s *orderService) EnqueueFCM(title, body string) {
	s.fcmCh <- FCMJob{Title: title, Body: body}
}

func (s *orderService) SaveToken(req dto.TokenRequest) error {
	s.tokensMu.Lock()
	defer s.tokensMu.Unlock()
	s.tokens[req.UserID] = req.Token
	return s.Repo.SaveToken(req)
}

func (s *orderService) BroadcastNow(req dto.FCMRequest) (dto.FCMResponse, error) {
	productJSON, err := json.Marshal(req.Product)
	if err != nil {
		return dto.FCMResponse{}, err
	}

	// Buat copy request dengan string product (opsional kalau mau simpan)
	reqWithJSON := req
	reqWithJSON.ProductJSON = string(productJSON)

	// Kirim ke repo
	created, err := s.Repo.SaveOrder(reqWithJSON)
	if err != nil {
		return dto.FCMResponse{}, err
	}

	tokens, err := s.Repo.GetAllTokens()
	if err != nil {
		return dto.FCMResponse{}, err
	}

	var lastRes string
	for _, token := range tokens {
		msg := &messaging.Message{
			Token: token,
			Notification: &messaging.Notification{
				Title: fmt.Sprintf("Web order: %s", created.Name),
				Body:  "You have a new order from the link, check now",
			},
		}

		res, err := s.client.Send(context.Background(), msg)
		if err != nil {
			log.Println("FCM send error:", err)
			_ = s.Repo.DeleteToken(token)
			continue
		}

		lastRes = res
	}

	return dto.FCMResponse{Name: lastRes}, nil
}

func (s *orderService) fcmWorker() {
	for job := range s.fcmCh {
		_, _ = s.BroadcastNow(dto.FCMRequest{
			Title: job.Title,
			Body:  job.Body,
		})
	}
}

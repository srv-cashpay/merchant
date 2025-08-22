package dashboard

import (
	"context"
	"log"

	"firebase.google.com/go/messaging"
)

func (s *dashboardService) EnqueueFCM(title, body string) {
	s.fcmCh <- FCMJob{Title: title, Body: body}
}

func (s *dashboardService) SaveToken(userID, token string) error {
	s.tokensMu.Lock()
	defer s.tokensMu.Unlock()
	s.tokens[userID] = token
	return s.Repo.SaveToken(userID, token)
}

func (s *dashboardService) fcmWorker() {
	for job := range s.fcmCh {
		tokens, err := s.Repo.GetAllTokens()
		if err != nil {
			log.Println("GetAllTokens error:", err)
			continue
		}

		for _, token := range tokens {
			msg := &messaging.Message{
				Token: token,
				Notification: &messaging.Notification{
					Title: job.Title, // contoh: "Cashpay Info"
					Body:  job.Body,  // contoh: "Saldo kamu bertambah Rp50.000"
				},
				Android: &messaging.AndroidConfig{
					Priority: "high",
				},
				APNS: &messaging.APNSConfig{
					Headers: map[string]string{"apns-priority": "10"},
				},
			}

			_, err := s.client.Send(context.Background(), msg)
			if err != nil {
				log.Println("FCM send error:", err)
				_ = s.Repo.DeleteToken(token)
			}
		}
	}
}

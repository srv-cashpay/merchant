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
	return nil
}

func (s *dashboardService) fcmWorker() {
	for job := range s.fcmCh {
		tokens, err := s.Repo.GetAllTokens()
		if err != nil {
			log.Println("GetAllTokens error:", err)
			continue
		}

		for _, token := range tokens {
			message := &messaging.Message{
				Token: token,
				Notification: &messaging.Notification{
					Title: job.Title,
					Body:  job.Body,
				},
			}

			_, err := s.client.Send(context.Background(), message)
			if err != nil {
				log.Println("FCM send error:", err)
				_ = s.Repo.DeleteToken(token) // hapus token invalid
			}
		}
	}
}

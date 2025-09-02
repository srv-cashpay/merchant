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

func (s *dashboardService) BroadcastNow(title, body string) ([]map[string]string, error) {
	tokens, err := s.Repo.GetAllTokens()
	if err != nil {
		return nil, err
	}

	var responses []map[string]string
	for _, token := range tokens {
		msg := &messaging.Message{
			Token: token,
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
		}

		res, err := s.client.Send(context.Background(), msg)
		if err != nil {
			log.Println("FCM send error:", err)
			_ = s.Repo.DeleteToken(token)
			continue
		}

		responses = append(responses, map[string]string{"name": res})
	}

	return responses, nil
}

func (s *dashboardService) fcmWorker() {
	for job := range s.fcmCh {
		_, _ = s.BroadcastNow(job.Title, job.Body)
	}
}

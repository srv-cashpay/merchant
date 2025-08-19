package dashboard

import (
	"context"
	"log"

	"firebase.google.com/go/messaging"
)

func (s *dashboardService) SaveToken(userID, token string) error {
	return s.Repo.SaveToken(userID, token)
}

func (s *dashboardService) BroadcastFCM(title, body string) error {
	tokens, err := s.Repo.GetAllTokens()
	if err != nil {
		return err
	}

	for _, t := range tokens {
		msg := &messaging.Message{
			Token: t,
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
		}

		_, err := s.client.Send(context.Background(), msg)
		if err != nil {
			log.Println("FCM send error:", err)
		}
	}

	return nil
}

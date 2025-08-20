package dashboard

import (
	"context"
	"log"
	"sync"

	"firebase.google.com/go/messaging"
)

func (s *dashboardService) SaveToken(userID, token string) error {
	return s.Repo.SaveToken(userID, token)
}

func (s *dashboardService) EnqueueFCM(title, body string) {
	s.fcmCh <- FCMJob{Title: title, Body: body}
}

func (s *dashboardService) startFCMWorkers(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range s.fcmCh {
				s.BroadcastFCM(job.Title, job.Body)
			}
		}()
	}
	wg.Wait()
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
			// hapus token invalid
			if delErr := s.Repo.DeleteToken(t); delErr != nil {
				log.Println("Delete token error:", delErr)
			}
		}
	}
	return nil
}

package dashboard

import (
	"context"
	"log"

	m "github.com/srv-cashpay/middlewares/middlewares"

	"firebase.google.com/go/messaging"

	firebase "firebase.google.com/go"

	"github.com/srv-cashpay/merchant/dto"
	r "github.com/srv-cashpay/merchant/repositories/dashboard"
	"google.golang.org/api/option"
)

type DashboardService interface {
	Get(req dto.GetDashboardRequest) (dto.GetDashboardResponse, error)
	SaveToken(userID, token string) error
	BroadcastFCM(title, body string) error
	EnqueueFCM(title, body string)
}

type FCMJob struct {
	Title string
	Body  string
}

type dashboardService struct {
	Repo   r.DomainRepository
	jwt    m.JWTService
	client *messaging.Client
	fcmCh  chan FCMJob
}

func NewDashboardService(Repo r.DomainRepository, jwtS m.JWTService) *dashboardService {
	credFile := "firebase-service-account.json"

	opt := option.WithCredentialsFile(credFile)
	// conf := &firebase.Config{ProjectID: "cashpay-2ac49"}

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln("Firebase init error:", err)
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalln("Firebase Messaging client error:", err)
	}

	s := &dashboardService{
		Repo:   Repo,
		jwt:    jwtS,
		client: client,
		fcmCh:  make(chan FCMJob, 100),
	}

	go s.startFCMWorkers(5) // 5 workers
	return s
}

// func NewDashboardService(Repo r.DomainRepository, jwtS m.JWTService) DashboardService {
// 	credFile := "firebase-service-account.json"

// 	opt := option.WithCredentialsFile(credFile)

// 	// isi project id sesuai JSON kamu
// 	conf := &firebase.Config{ProjectID: "cashpay-2ac49"}

// 	app, err := firebase.NewApp(context.Background(), conf, opt)
// 	if err != nil {
// 		log.Fatalf("error init firebase: %v", err)
// 	}

// 	client, err := app.Messaging(context.Background())
// 	if err != nil {
// 		log.Fatalf("error init fcm client: %v", err)
// 	}

// 	return &dashboardService{
// 		Repo:   Repo,
// 		jwt:    jwtS,
// 		client: client,
// 	}
// }

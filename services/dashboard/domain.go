package dashboard

import (
	"context"
	"log"
	"sync"

	m "github.com/srv-cashpay/middlewares/middlewares"

	"firebase.google.com/go/messaging"

	firebase "firebase.google.com/go"

	"github.com/srv-cashpay/merchant/dto"
	r "github.com/srv-cashpay/merchant/repositories/dashboard"
	"google.golang.org/api/option"
)

type DashboardService interface {
	Get(req dto.GetDashboardRequest) (dto.GetDashboardResponse, error)
	EnqueueFCM(title, body string)
	SaveToken(userID, token string) error
}

type FCMJob struct {
	Title string
	Body  string
}

type dashboardService struct {
	Repo     r.DomainRepository
	jwt      m.JWTService
	client   *messaging.Client
	fcmCh    chan FCMJob
	wg       sync.WaitGroup
	tokens   map[string]string
	tokensMu sync.Mutex
}

func NewDashboardService(Repo r.DomainRepository, jwtS m.JWTService) *dashboardService {

	credFile := "firebase-service-account.json"
	opt := option.WithCredentialsFile(credFile)
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
		client: client,
		fcmCh:  make(chan FCMJob, 100),
	}

	// start worker pool
	for i := 0; i < 5; i++ {
		go s.fcmWorker()
	}

	return s
}

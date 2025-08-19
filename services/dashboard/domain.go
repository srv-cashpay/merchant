package dashboard

import (
	"context"
	"log"

	m "github.com/srv-cashpay/middlewares/middlewares"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/srv-cashpay/merchant/dto"
	r "github.com/srv-cashpay/merchant/repositories/dashboard"
	"google.golang.org/api/option"
)

type DashboardService interface {
	Get(req dto.GetDashboardRequest) (dto.GetDashboardResponse, error)
	SaveToken(userID, token string) error
	BroadcastFCM(title, body string) error
}

type dashboardService struct {
	Repo   r.DomainRepository
	jwt    m.JWTService
	client *messaging.Client
}

func NewDashboardService(Repo r.DomainRepository, jwtS m.JWTService) DashboardService {
	credFile := "/app/configs/firebase-service-account.json"

	opt := option.WithCredentialsFile(credFile)

	// isi project id sesuai JSON kamu
	conf := &firebase.Config{ProjectID: "cashpay-2ac49"}

	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalf("error init firebase: %v", err)
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error init fcm client: %v", err)
	}

	return &dashboardService{
		Repo:   Repo,
		jwt:    jwtS,
		client: client,
	}
}

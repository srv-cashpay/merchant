package order

import (
	"context"
	"log"
	"os"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/merchant/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"
	"google.golang.org/api/option"

	r "github.com/srv-cashpay/merchant/repositories/order"
)

type OrderService interface {
	Create(req dto.OrderRequest) (dto.OrderResponse, error)
	Order(req dto.OrderRequest) (dto.OrderResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Delete(req dto.DeleteOrderRequest) (dto.DeleteOrderResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdOrderRequest) (*dto.OrderResponse, error)
	Update(req dto.OrderUpdateRequest) (dto.OrderUpdateResponse, error)

	EnqueueFCM(title, body string)
	SaveToken(dto.TokenRequest) error
	BroadcastNow(req dto.FCMRequest) (dto.FCMResponse, error)
}

type orderService struct {
	Repo     r.DomainRepository
	jwt      m.JWTService
	client   *messaging.Client
	fcmCh    chan FCMJob
	wg       sync.WaitGroup
	tokens   map[string]string
	tokensMu sync.Mutex
}

type FCMJob struct {
	Title string
	Body  string
}

func NewOrderService(Repo r.DomainRepository, jwtS m.JWTService) OrderService {
	credFile := os.Getenv("CredFile")
	opt := option.WithCredentialsFile(credFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln("Firebase init error:", err)
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalln("Firebase Messaging client error:", err)
	}

	s := &orderService{
		Repo:   Repo,
		jwt:    jwtS,
		client: client,
		fcmCh:  make(chan FCMJob, 100),
		tokens: make(map[string]string), // ✅ inisialisasi map
	}
	for i := 0; i < 5; i++ {
		go s.fcmWorker()
	}
	return s
}

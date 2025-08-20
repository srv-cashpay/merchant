package dashboard

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	s "github.com/srv-cashpay/merchant/services/dashboard"
)

type DomainHandler interface {
	Get(c echo.Context) error
	HandleWebSocket(c echo.Context) error
	SaveToken(c echo.Context) error
	SendBroadcast(c echo.Context) error
}

type domainHandler struct {
	serviceDashboard s.DashboardService
	clients          map[*websocket.Conn]bool
	broadcast        chan []byte
	mu               sync.Mutex
}

func NewDashboardHandler(service s.DashboardService) DomainHandler {
	return &domainHandler{
		serviceDashboard: service,
		clients:          make(map[*websocket.Conn]bool),
		broadcast:        make(chan []byte, 100), // buffered
	}
}

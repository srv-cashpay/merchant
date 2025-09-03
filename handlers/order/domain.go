package order

import (
	"sync"

	"github.com/gorilla/websocket"
	s "github.com/srv-cashpay/merchant/services/order"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	Order(c echo.Context) error
	Get(c echo.Context) error
	GetById(c echo.Context) error
	BulkDelete(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error

	HandleWebSocket(c echo.Context) error
	SaveToken(c echo.Context) error
	SendBroadcast(c echo.Context) error
}

type domainHandler struct {
	serviceOrder s.OrderService
	clients      map[*websocket.Conn]bool
	broadcast    chan []byte
	mu           sync.Mutex
}

func NewOrderHandler(service s.OrderService) DomainHandler {
	return &domainHandler{
		serviceOrder: service,
		clients:      make(map[*websocket.Conn]bool),
		broadcast:    make(chan []byte, 100),
	}
}

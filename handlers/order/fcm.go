package order

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/merchant/dto"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (h *domainHandler) HandleWebSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return err
	}

	h.mu.Lock()
	h.clients[conn] = true
	h.mu.Unlock()
	log.Println("Client connected:", conn.RemoteAddr())

	go h.readPump(conn)
	return nil
}

func (h *domainHandler) readPump(conn *websocket.Conn) {
	defer func() {
		h.mu.Lock()
		delete(h.clients, conn)
		h.mu.Unlock()
		conn.Close()
		log.Println("Client disconnected:", conn.RemoteAddr())
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		h.broadcast <- msg

		// setiap pesan dari WS → enqueue FCM
		h.serviceOrder.EnqueueFCM("Pesan Baru", string(msg))
	}
}

func (h *domainHandler) SendBroadcast(c echo.Context) error {
	var req dto.FCMRequest
	var res dto.FCMResponse

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	res, err := h.serviceOrder.BroadcastNow(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if res.Name == "" {
		return c.JSON(http.StatusOK, map[string]string{"status": "no tokens"})
	}

	return c.JSON(http.StatusOK, map[string]string{"name": res.Name})
}

func (h *domainHandler) SaveToken(c echo.Context) error {
	var req dto.TokenRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.serviceOrder.SaveToken(req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, dto.TokenResponse{Status: "ok"})
}

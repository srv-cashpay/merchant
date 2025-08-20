package dashboard

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
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

	conn.SetReadLimit(1024)
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		h.broadcast <- msg
		// kirim FCM via worker pool
		go h.serviceDashboard.EnqueueFCM("Pesan Baru", string(msg))
	}
}

func (h *domainHandler) StartBroadcaster() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case msg := <-h.broadcast:
			h.mu.Lock()
			for conn := range h.clients {
				conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
				if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					log.Println("Write error:", err)
					conn.Close()
					delete(h.clients, conn)
				}
			}
			h.mu.Unlock()
		case <-ticker.C:
			// ping semua client agar tetap alive
			h.mu.Lock()
			for conn := range h.clients {
				conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
				if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					conn.Close()
					delete(h.clients, conn)
				}
			}
			h.mu.Unlock()
		}
	}
}

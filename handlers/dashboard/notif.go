package dashboard

import (
	"log"
	"net/http"
	"sync"

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

var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex
var broadcast = make(chan []byte)

// HandleWebSocket → semua client connect ke sini
func (h *domainHandler) HandleWebSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return err
	}

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	log.Println("Client connected:", conn.RemoteAddr())

	go func(conn *websocket.Conn) {
		defer func() {
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()
			conn.Close()
			log.Println("Client disconnected:", conn.RemoteAddr())
		}()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}
			log.Printf("Received: %s", msg)

			// kirim ke semua client websocket
			broadcast <- append([]byte{}, msg...)

			// kirim juga ke semua device via FCM
			go h.serviceDashboard.BroadcastFCM("Pesan Baru", string(msg))
		}
	}(conn)

	return nil
}

// Broadcaster kirim pesan ke semua client
func StartBroadcaster() {
	for {
		msg := <-broadcast
		mu.Lock()
		for conn := range clients {
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("Write error:", err)
				conn.Close()
				delete(clients, conn)
			}
		}
		mu.Unlock()
	}
}

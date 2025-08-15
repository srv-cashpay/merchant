package dashboard

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

// Konfigurasi upgrader WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Untuk testing, izinkan semua origin
		// TODO: Batasi origin untuk production
		return true
	},
}

// HandleWebSocket adalah handler untuk koneksi WebSocket
func (b *domainHandler) HandleWebSocket(c echo.Context) error {
	// Gunakan c.Response().Writer untuk Upgrade
	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return err
	}
	defer func() {
		log.Println("Client disconnected:", conn.RemoteAddr())
		conn.Close()
	}()

	log.Println("Client connected:", conn.RemoteAddr())

	for {
		// Membaca pesan dari client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}
		log.Printf("Received: %s", message)

		// Kirim balik pesan ke client
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("Write Error:", err)
			break
		}
	}

	// Penting: return nil supaya Echo tidak kirim HTTP response
	return nil
}

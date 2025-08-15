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
		return true // <- Izinkan semua origin (untuk testing)
	},
}

// Handler WebSocket
func (b *domainHandler) HandleWebSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return err
	}
	defer conn.Close()

	log.Println("Client connected:", conn.RemoteAddr())

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}
		log.Printf("Received: %s", message)

		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("Write Error:", err)
			break
		}
	}

	return nil
}

// func main() {
// 	e := echo.New()

// 	// Route untuk WebSocket
// 	e.GET("/ws", handleWebSocket)

// 	host := "0.0.0.0:8080"
// 	fmt.Println("WebSocket server started at http://" + host)

// 	if err := e.Start(host); err != nil {
// 		log.Fatal("Echo Start Error:", err)
// 	}
// }

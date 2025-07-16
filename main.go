package main

import (
	"log"

	"github.com/olivere/elastic/v7"
	"github.com/srv-cashpay/merchant/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Client *elastic.Client

func main() {
	var error error

	Client, error = elastic.NewClient(
		elastic.SetURL("http://elasticsearch:9200"), // jika pakai docker service
		elastic.SetSniff(false),
	)
	if error != nil {
		log.Fatalf("Failed to create Elastic client: %v", error)
	}

	e := routes.New()

	e.Use(middleware.CORS())

	// Sertifikat Let's Encrypt
	certFile := "/certs/fullchain.pem"
	keyFile := "/certs/privkey.pem"

	// Jalankan HTTPS langsung dari Echo
	err := e.StartTLS(":2345", certFile, keyFile)
	if err != nil {
		log.Fatal("StartTLS error: ", err)
	}
}

// CORSMiddleware ..
func CORSMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(204)
			}

			return next(c)
		}
	}
}

package authenticator_request

import (
	s "github.com/srv-cashpay/merchant/services/authenticator_request"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
}

type domainHandler struct {
	serviceAuthenticator s.AuthenticatorService
}

func NewAuthenticatorHandler(service s.AuthenticatorService) DomainHandler {
	return &domainHandler{
		serviceAuthenticator: service,
	}
}

package routing

import (
	"github.com/dannyvelas/go-backend/auth"
)

type RoutingAuthenticator struct {
	auth.Authenticator
}

func NewAuthenticator(authenticator auth.Authenticator) RoutingAuthenticator {
	return RoutingAuthenticator{
		Authenticator: authenticator,
	}
}

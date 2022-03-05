package auth

import (
	"github.com/dannyvelas/go-backend/config"
)

type Authenticator struct {
	tokenSecret []byte
}

func NewAuthenticator(tokenConfig config.TokenConfig) Authenticator {
	return Authenticator{tokenSecret: []byte(tokenConfig.Secret())}
}

package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/dannyvelas/examplego_api/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

var (
	errNotSigningMethodHMAC = errors.New("jwt: Not using SigningMethodHMAC")
	errCastingJWTClaims     = errors.New("jwt: Failed to cast JWT token to JWTClaims struct")
	errInvalidToken         = errors.New("jwt: Invalid Token")
)

type jwtClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

type JWTMiddleware struct {
	tokenSecret []byte
}

func NewJWTMiddleware(tokenConfig config.TokenConfig) JWTMiddleware {
	return JWTMiddleware{tokenSecret: []byte(tokenConfig.Secret())}
}

func (jwtMiddleware JWTMiddleware) newJWT(id string) (string, error) {
	claims := jwtClaims{
		id,
		jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * 15).Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtMiddleware.tokenSecret)
}

func (jwtMiddleware JWTMiddleware) parseJWT(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errNotSigningMethodHMAC
		}

		return jwtMiddleware.tokenSecret, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*jwtClaims); !ok || !token.Valid {
		if !ok {
			return "", errCastingJWTClaims
		} else {
			return "", errInvalidToken
		}
	} else {
		return claims.Id, nil
	}
}

func (jwtMiddleware JWTMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("JWT Middleware")

		cookie, err := r.Cookie("jwt")
		if err != nil {
			err = fmt.Errorf("jwt_middleware: Rejected Auth: %v", err)
			respondError(w, err, errUnauthorized)
			return
		}

		userId, err := jwtMiddleware.parseJWT(cookie.Value)
		if err != nil {
			err = fmt.Errorf("jwt_middleware: Rejected Auth: Error parsing jwt cookie: %v", err)
			respondError(w, err, errUnauthorized)
			return
		}

		ctx := r.Context()
		updatedCtx := context.WithValue(ctx, "id", userId)
		updatedReq := r.WithContext(updatedCtx)

		next.ServeHTTP(w, updatedReq)
	})
}

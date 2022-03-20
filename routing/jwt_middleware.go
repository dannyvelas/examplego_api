package routing

import (
	"context"
	"errors"
	"fmt"
	"github.com/dannyvelas/examplego_api/apierror"
	"github.com/dannyvelas/examplego_api/config"
	"github.com/dannyvelas/examplego_api/routing/internal"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type JWTMiddleware struct {
	tokenSecret []byte
}

func NewJWTMiddleware(tokenConfig config.TokenConfig) JWTMiddleware {
	return JWTMiddleware{tokenSecret: []byte(tokenConfig.Secret())}
}

type JWTClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func (jwtMiddleware JWTMiddleware) newJWT(id string) (string, error) {
	claims := JWTClaims{
		id,
		jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * 15).Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtMiddleware.tokenSecret)
}

func (jwtMiddleware JWTMiddleware) parseJWT(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Not using SigningMethodHMAC")
		}

		return jwtMiddleware.tokenSecret, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*JWTClaims); !ok || !token.Valid {
		if !ok {
			return "", errors.New("Failure casting JWTClaims")
		} else {
			return "", errors.New("Token not valid")
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
			err = fmt.Errorf("Rejected Authorization: %v", err)
			internal.RespondError(w, err, apierror.Unauthorized)
			return
		}

		userId, err := jwtMiddleware.parseJWT(cookie.Value)
		if err != nil {
			err = fmt.Errorf("Rejected Authorization: Error parsing jwt cookie: %v", err)
			internal.RespondError(w, err, apierror.Unauthorized)
			return
		}

		ctx := r.Context()
		updatedCtx := context.WithValue(ctx, "id", userId)
		updatedReq := r.WithContext(updatedCtx)

		next.ServeHTTP(w, updatedReq)
	})
}

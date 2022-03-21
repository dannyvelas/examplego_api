package routing

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dannyvelas/examplego_api/routing/internal"
	"github.com/dannyvelas/examplego_api/storage"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type credentials struct {
	Id       string
	Password string
}

func Login(jwtMiddleware JWTMiddleware, adminsRepo storage.AdminsRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Login Endpoint")

		var creds credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			err = fmt.Errorf("login_router: Error decoding credentials body: %v", err)
			internal.RespondError(w, err, internal.ErrBadRequest)
			return
		}

		admin, err := adminsRepo.GetOne(creds.Id)
		if errors.Is(err, storage.ErrNoRows) {
			err = fmt.Errorf("login_router: Rejected Auth: %v", err)
			internal.RespondError(w, err, internal.ErrUnauthorized)
			return
		} else if err != nil {
			err = fmt.Errorf("login_router: Error querying adminsRepo: %v", err)
			internal.RespondError(w, err, internal.ErrInternalServerError)
			return
		}

		if err = bcrypt.CompareHashAndPassword(
			[]byte(admin.Password),
			[]byte(creds.Password),
		); err != nil {
			err = fmt.Errorf("login_router: Rejected Auth: %v", err)
			internal.RespondError(w, err, internal.ErrUnauthorized)
			return
		}

		token, err := jwtMiddleware.newJWT(admin.Id)
		if err != nil {
			err = fmt.Errorf("login_router: Error generating JWT: %v", err)
			internal.RespondError(w, err, internal.ErrInternalServerError)
			return
		}

		cookie := http.Cookie{Name: "jwt", Value: token, HttpOnly: true, Path: "/"}
		http.SetCookie(w, &cookie)
	}
}

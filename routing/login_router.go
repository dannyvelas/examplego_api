package routing

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dannyvelas/examplego_api/apierror"
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

func Login(jwtMiddleware JWTMiddleware, adminRepo storage.AdminRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Login Endpoint")

		var creds credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			err = fmt.Errorf("Error decoding credentials body: %q", err)
			err = apierror.WrapSentinel(err, apierror.ErrBadRequest)
			internal.RespondError(w, err)
			return
		}

		admin, err := adminRepo.GetOne(creds.Id)
		if errors.Is(err, apierror.ErrNotFound) {
			err = apierror.WrapSentinel(err, apierror.ErrUnauthorized)
			internal.RespondError(w, err)
			return
		} else if err != nil {
			err = fmt.Errorf("Error in adminRepo.GetOne: %q", err)
			internal.RespondError(w, err)
			return
		}

		if err = bcrypt.CompareHashAndPassword(
			[]byte(admin.Password),
			[]byte(creds.Password),
		); err != nil {
			err = apierror.WrapSentinel(err, apierror.ErrUnauthorized)
			internal.RespondError(w, err)
			return
		}

		token, err := jwtMiddleware.newJWT(admin.Id)
		if err != nil {
			err = fmt.Errorf("Error generating JWT: %q", err)
			internal.RespondError(w, err)
			return
		}

		cookie := http.Cookie{Name: "jwt", Value: token, HttpOnly: true, Path: "/"}
		http.SetCookie(w, &cookie)
	}
}

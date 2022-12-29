package router

import (
	"net/http"

	"github.com/nozomi-iida/nozo_blog/presentation/controller/auth"
)

type authRouter struct {
	ac auth.AuthController
}

func NewRouter(fileString string) (authRouter, error)  {
	ac, err := auth.NewAuthController(fileString)
	if err != nil {
		return authRouter{}, err
	}
	return authRouter{ac: ac}, nil
}

func (ar *authRouter) HandleSignUpRequest(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "POST":
		ar.ac.SignUpRequest(w, r)
	default:
		w.WriteHeader(405)
	}
}

func (ar *authRouter) HandleSignInRequest(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "POST":
		ar.ac.SignInRequest(w, r)
	default:
		w.WriteHeader(405)
	}
}

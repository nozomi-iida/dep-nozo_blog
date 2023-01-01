package router

import (
	"net/http"

	"github.com/nozomi-iida/nozo_blog/middleware"
	"github.com/nozomi-iida/nozo_blog/presentation"
	"github.com/nozomi-iida/nozo_blog/presentation/controller"
)

type authRouter struct {
	atc controller.AuthController
	ac controller.ArticleController
}

func NewRouter(fileString string) (authRouter, error)  {
	atc, err := controller.NewAuthController(fileString)
	ac, err := controller.NewArticleController(fileString)
	if err != nil {
		return authRouter{}, err
	}
	return authRouter{atc: atc, ac: ac}, nil
}

func (ar *authRouter) HandleSignUpRequest(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "POST":
		ar.atc.SignUpRequest(w, r)
	default:
		presentation.ErrorHandler(w, presentation.ErrStatusMethodNotAllowed)
	}
}

func (ar *authRouter) HandleSignInRequest(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "POST":
		ar.atc.SignInRequest(w, r)
	default:
		presentation.ErrorHandler(w, presentation.ErrStatusMethodNotAllowed)
	}
}

func (ar *authRouter) HandleArticleRequest(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case http.MethodPost:
		middleware.AuthMiddleware(http.HandlerFunc(ar.ac.PostRequest))	
	default:
		presentation.ErrorHandler(w, presentation.ErrStatusMethodNotAllowed)
	}
}

package presentation

import (
	"net/http"

	"github.com/nozomi-iida/nozo_blog/presentation/controller"
	"github.com/nozomi-iida/nozo_blog/presentation/helpers"
	"github.com/nozomi-iida/nozo_blog/presentation/middleware"
)

type router struct {
	atc controller.AuthController
	ac controller.ArticleController
	tc controller.TopicController
}

func NewRouter(fileString string) (router, error)  {
	atc, err := controller.NewAuthController(fileString)
	ac, err := controller.NewArticleController(fileString)
	tc, err := controller.NewTopicController(fileString)
	if err != nil {
		return router{}, err
	}
	return router{atc: atc, ac: ac, tc: tc}, nil
}

func (rt *router) HandleSignUpRequest(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case http.MethodPost:
		rt.atc.SignUpRequest(w, r)
	default:
		helpers.ErrorHandler(w, helpers.ErrStatusMethodNotAllowed)
	}
}

func (rt *router) HandleSignInRequest(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case http.MethodPost:
		rt.atc.SignInRequest(w, r)
	default:
		helpers.ErrorHandler(w, helpers.ErrStatusMethodNotAllowed)
	}
}

func (rt *router) HandleArticleRequest(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case http.MethodPost:
		middleware.AuthMiddleware(http.HandlerFunc(rt.ac.PostRequest))	
	default:
		helpers.ErrorHandler(w, helpers.ErrStatusMethodNotAllowed)
	}
}

func (rt *router) HandleTopicRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		middleware.AuthMiddleware(http.HandlerFunc(rt.tc.CreteRequest))
	default:
		helpers.ErrorHandler(w, helpers.ErrStatusMethodNotAllowed)
	}
} 

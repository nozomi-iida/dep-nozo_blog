package presentation

import (
	"net/http"
	"path"
	"strings"

	"github.com/nozomi-iida/nozo_blog/presentation/controller"
	"github.com/nozomi-iida/nozo_blog/presentation/helpers"
	"github.com/nozomi-iida/nozo_blog/presentation/middleware"
)

type router struct {
	atc controller.AuthController
	ac controller.ArticleController
	tc controller.TopicController
}

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
			return p[1:], "/"
	}
	return p[1:i], p[i:]
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

// routesの構造化ってどうやるんだ？ findByIDとListが同じエンドポイントになってしまっている
func (rt *router) HandleArticleRequest(w http.ResponseWriter, r *http.Request)  {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	head, r.URL.Path = shiftPath(r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		if head != "" {
			http.HandlerFunc(rt.ac.FindByIdRequest).ServeHTTP(w, r)
		} else {
			http.HandlerFunc(rt.ac.ListRequest).ServeHTTP(w, r)
		}
	case http.MethodPost:
		if head != "" {
			middleware.AuthMiddleware(http.HandlerFunc(rt.ac.PostRequest)).ServeHTTP(w, r)
		}
	case http.MethodPatch:
		if head != "" {
			middleware.AuthMiddleware(http.HandlerFunc(rt.ac.PatchRequest)).ServeHTTP(w, r)
		}
	case http.MethodDelete:
		if head != "" {
			middleware.AuthMiddleware(http.HandlerFunc(rt.ac.DeleteRequest)).ServeHTTP(w, r)
		}
	default:
		helpers.ErrorHandler(w, helpers.ErrStatusMethodNotAllowed)
	}
}

func (rt *router) HandleTopicRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.HandlerFunc(rt.tc.ListRequest).ServeHTTP(w, r)
	case http.MethodPost:
		middleware.AuthMiddleware(http.HandlerFunc(rt.tc.CreteRequest)).ServeHTTP(w, r)
	default:
		helpers.ErrorHandler(w, helpers.ErrStatusMethodNotAllowed)
	}
} 

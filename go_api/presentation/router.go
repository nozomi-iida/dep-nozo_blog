package presentation

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nozomi-iida/nozo_blog_go_api/presentation/controller"
	admincontroller "github.com/nozomi-iida/nozo_blog_go_api/presentation/controller/admin-controller"
	"github.com/nozomi-iida/nozo_blog_go_api/presentation/middleware"
)

func NewRouter(fileString string) (*chi.Mux, error) {
	atc, err := controller.NewAuthController(fileString)
	ac, err := controller.NewArticleController(fileString)
	tc, err := controller.NewTopicController(fileString)
	tgc, err := controller.NewTagController(fileString)
	aac, err := admincontroller.NewArticleController(fileString)
	adtc, err := admincontroller.NewTopicController(fileString)

	if err != nil {
		return &chi.Mux{}, err
	}
	r := chi.NewRouter()
	r.Use(middleware.WrapHandlerWithLoggingMiddleware)
	r.Use(middleware.CorsMiddleware)
	r.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
	})
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign_in", atc.SignInRequest)
			r.Post("/sign_up", atc.SignUpRequest)
		})
		r.Route("/public", func(r chi.Router) {
			r.Route("/articles", func(r chi.Router) {
				r.Get("/", ac.ListRequest)
				r.Get("/{article_id}", ac.FindByIdRequest)
			})
			r.Route("/topics", func(r chi.Router) {
				r.Get("/", tc.ListRequest)
				r.Get("/{name}", tc.FindByNameRequest)
			})
			r.Route("/tags", func(r chi.Router) {
				r.Get("/", tgc.ListRequest)
			})
		})
		r.Route("/admin", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Route("/articles", func(r chi.Router) {
				r.Get("/", aac.ListRequest)
				r.Post("/", ac.PostRequest)
				r.Get("/{article_id}", aac.FindByIdRequest)
				r.Patch("/{article_id}", aac.PatchRequest)
				r.Delete("/{article_id}", ac.DeleteRequest)
			})
			r.Route("/topics", func(r chi.Router) {
				r.Get("/", adtc.ListRequest)
				r.Post("/", adtc.PostRequest)
				r.Patch("/{topic_id}", adtc.PatchRequest)
				// r.Delete("/{topic_id}", adtc.DeleteRequest)
			})
		})
	})

	return r, nil
}

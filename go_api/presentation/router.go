package presentation

import (
	"github.com/go-chi/chi/v5"
	"github.com/nozomi-iida/nozo_blog/presentation/controller"
	"github.com/nozomi-iida/nozo_blog/presentation/middleware"
)

func NewRouter(fileString string) (*chi.Mux, error)  {
	atc, err := controller.NewAuthController(fileString)
	ac, err := controller.NewArticleController(fileString)
	tc, err := controller.NewTopicController(fileString)
	tgc, err := controller.NewTagController(fileString)
	if err != nil {
		return &chi.Mux{}, err
	}
	r := chi.NewRouter()
	r.Use(middleware.WrapHandlerWithLoggingMiddleware)
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign-in", atc.SignInRequest)
			r.Post("/sign-up", atc.SignUpRequest)
		})
		r.Route("/public", func(r chi.Router) {
			r.Route("/articles", func(r chi.Router) {
				r.Get("/", ac.ListRequest)
				r.Get("/{article_id}", ac.FindByIdRequest)
			})
			r.Route("/topics", func(r chi.Router) {
				r.Get("/", tc.ListRequest)
			})
			r.Route("/tags", func(r chi.Router) {
				r.Get("/", tgc.ListRequest)
			})
		})
		r.Route("/admin", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Route("/articles", func(r chi.Router) {
				r.Get("/", ac.ListRequest)
				r.Post("/", ac.PostRequest)
				r.Get("/{article_id}", ac.FindByIdRequest)
				r.Put("/{article_id}", ac.PatchRequest)
				r.Delete("/{article_id}", ac.DeleteRequest)
			})
		})
	})

	
	return r, nil
}

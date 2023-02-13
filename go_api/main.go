package main

import (
	"net/http"

	"github.com/nozomi-iida/nozo_blog/libs"
	"github.com/nozomi-iida/nozo_blog/presentation"
	"github.com/nozomi-iida/nozo_blog/presentation/middleware"
	"github.com/rs/cors"
)

var ar, _ = presentation.NewRouter("./tmp/data.db")


func logHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, middleware.WrapHandlerWithLoggingMiddleware(http.HandlerFunc(handler)).ServeHTTP)
}

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/sign_up", middleware.WrapHandlerWithLoggingMiddleware(http.HandlerFunc(ar.HandleSignUpRequest)).ServeHTTP)
	mux.HandleFunc("/sign_in", middleware.WrapHandlerWithLoggingMiddleware(http.HandlerFunc(ar.HandleSignInRequest)).ServeHTTP)
	mux.HandleFunc("/articles", middleware.WrapHandlerWithLoggingMiddleware(http.HandlerFunc(ar.HandleArticleRequest)).ServeHTTP)
	mux.HandleFunc("/topics", middleware.WrapHandlerWithLoggingMiddleware(http.HandlerFunc(ar.HandleTopicRequest)).ServeHTTP)
	// handler := middleware.CorsMiddleware(mux)
	handler := cors.Default().Handler(mux)
	libs.ZipLogger().Error(http.ListenAndServe(":8080", handler).Error())
}

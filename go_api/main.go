package main

import (
	"net/http"

	"github.com/nozomi-iida/nozo_blog/libs"
	"github.com/nozomi-iida/nozo_blog/presentation"
	"github.com/nozomi-iida/nozo_blog/presentation/middleware"
)

var ar, _ = presentation.NewRouter("./tmp/data.db")


func logHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, middleware.WrapHandlerWithLoggingMiddleware(http.HandlerFunc(handler)).ServeHTTP)
}

func main()  {
	logHandleFunc("/sign_up", ar.HandleSignUpRequest)
	logHandleFunc("/sign_in", ar.HandleSignInRequest)
	logHandleFunc("/articles", ar.HandleArticleRequest)
	logHandleFunc("/topics", ar.HandleTopicRequest)
	libs.ZipLogger().Error(http.ListenAndServe(":8080", nil).Error())
}

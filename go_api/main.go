package main

import (
	"log"
	"net/http"

	"github.com/nozomi-iida/nozo_blog/presentation"
)

var ar, _ = presentation.NewRouter("./tmp/data.db")

func main()  {
	http.HandleFunc("/sign_up", ar.HandleSignUpRequest)
	http.HandleFunc("/sign_in", ar.HandleSignInRequest)
	http.HandleFunc("/articles", ar.HandleArticleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

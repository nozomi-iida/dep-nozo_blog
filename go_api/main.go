package main

import (
	"log"
	"net/http"

	"github.com/nozomi-iida/nozo_blog/router"
)

var ar, _ = router.NewRouter("./tmp/data.db")

func main()  {
	http.HandleFunc("/sign_up", ar.HandleSignUpRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

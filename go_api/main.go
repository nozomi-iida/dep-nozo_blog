package main

import (
	"net/http"

	"github.com/nozomi-iida/nozo_blog_go_api/libs"
	"github.com/nozomi-iida/nozo_blog_go_api/presentation"
)

func main() {
	var r, _ = presentation.NewRouter("./tmp/data.db")
	// handler := cors.AllowAll().Handler(r)

	libs.ZipLogger().Error(http.ListenAndServe(":8080", r).Error())
}

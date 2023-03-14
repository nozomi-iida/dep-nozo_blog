package main

import (
	"net/http"

	"github.com/nozomi-iida/nozo_blog/libs"
	"github.com/nozomi-iida/nozo_blog/presentation"
)



func main()  {
	var r, _ = presentation.NewRouter("./tmp/data.db")
	libs.ZipLogger().Error(http.ListenAndServe(":8080", r).Error())
}

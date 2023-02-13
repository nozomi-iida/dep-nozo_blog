package middleware

import (
	"net/http"
)

// これでcors errorが出る理由が分からない
// `next.ServeHTTP(w, r)`を書かないとcors errorにはならない
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := w.Header()
		headers.Set("Access-Control-Allow-Origin", "*")
		headers.Set("Access-Control-Allow-Credentials", "true")
		headers.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		next.ServeHTTP(w, r)
	})	
}

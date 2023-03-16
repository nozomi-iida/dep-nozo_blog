package middleware

import (
	"net/http"
)

// これでcors errorが出る理由が分からない
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := w.Header()
		headers.Set("Access-Control-Allow-Origin", "*")
		headers.Set("Access-Control-Allow-Credentials", "true")
		headers.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT, PATCH")
		headers.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		headers.Set("ExposedHeaders", "Link")

		if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
		}

		next.ServeHTTP(w, r)	
	})	
}

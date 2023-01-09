package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/nozomi-iida/nozo_blog/presentation/helpers"
	"github.com/nozomi-iida/nozo_blog/valueobject"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			helpers.ErrorHandler(w, helpers.ErrUnauthorized)
		} else {
			jwtToken := authHeader[1]
			if jwtToken == "" {
				helpers.ErrorHandler(w, helpers.ErrUnauthorized)
				return
			}
			claims, err := valueobject.Decode(jwtToken) 
			if err != nil || claims.UserId.ID() == 0 {
				helpers.ErrorHandler(w, helpers.ErrUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "userId", claims.UserId)
		  next.ServeHTTP(w, r.WithContext(ctx))
		}
	})	
}

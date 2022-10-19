package auth

import (
	"context"
	"net/http"
)

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctxKey := "X-Api-Key"
			ctxValue := r.Header.Get("X-Api-Key")

			// put it in context
			ctx := context.WithValue(r.Context(), ctxKey, ctxValue)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

package middlewares

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	service "github.com/jn0x/reddigo/services"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "applicaition/json")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token missing"))
			return
		}

		token := service.VerifyToken(tokenString)
		if token == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "claims", token.Claims.(jwt.MapClaims))
		if token.Valid {

			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid  token"))
	})
}

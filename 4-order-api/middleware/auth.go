package middleware

import (
	"context"
	"go/hw/4-order-api/configs"
	"go/hw/4-order-api/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	contextPhoneKey key = "contextPhoneKey"
)

func AuthMiddleware(next http.Handler, cfg *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, claims := jwt.NewJWT(cfg.JWT.Secret).VerifyToken(token)
		if !isValid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), contextPhoneKey, claims.Phone)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

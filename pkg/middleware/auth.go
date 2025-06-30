package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/GrudTrigger/trainin_tracker/configs"
	"github.com/GrudTrigger/trainin_tracker/pkg/jwt"
)


type contextKey string

const UserContextKey contextKey = "user"

func IsAuthed(next http.Handler, cfg *configs.Configs) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")

		if authedHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		token := strings.TrimPrefix(authedHeader, "Bearer ")
		if token == "" {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}

		isValid, data := jwt.NewJwt(cfg.Secret).Parse(token)
		if !isValid || data == nil {
			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ForContext(ctx context.Context) *jwt.JWTData {
	raw, _ := ctx.Value(UserContextKey).(*jwt.JWTData)
	return raw
}

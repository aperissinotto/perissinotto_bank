package middleware

import (
	"context"
	"net/http"

	"github.com/aperissinotto/perissinotto_bank/internal/infrastructure/auth"
	"github.com/aperissinotto/perissinotto_bank/internal/interfaces/http/handler"
)

type contextKey string

const ClienteContextKey contextKey = "cliente"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session_token")
		if err != nil {
			handler.WriteError(
				w,
				http.StatusUnauthorized,
				"Não autenticado!",
				err.Error(),
			)
			return
		}

		claims, err := auth.ValidarToken(cookie.Value)
		if err != nil {
			handler.WriteError(
				w,
				http.StatusUnauthorized,
				"Sessão inválida!",
				err.Error(),
			)
			return
		}

		ctx := context.WithValue(r.Context(), ClienteContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

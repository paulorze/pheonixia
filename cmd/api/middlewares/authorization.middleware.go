package middlewares

import (
	"context"
	"net/http"
	"phoenixia/utils"
	"strings"
)

type ContextKey string

const (
	UserIDKey    ContextKey = "userID"
	UsernameKey  ContextKey = "username"
	RoleKey      ContextKey = "role"
	DocumentsKey ContextKey = "documents"
)

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Eliminar el prefijo "Bearer " si está presente
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Validar el token y extraer los claims
		id, username, role, documents, err := utils.ValidateJWTSession(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Colocar el userID en el contexto usando la clave personalizada
		ctx := context.WithValue(r.Context(), UserIDKey, id)
		ctx = context.WithValue(ctx, UsernameKey, username)
		ctx = context.WithValue(ctx, RoleKey, role)
		ctx = context.WithValue(ctx, DocumentsKey, documents)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

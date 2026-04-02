package middleware

import (
	"context"
	"net/http"

	"github.com/devlucas-java/lucatask/internal/infra/jwt"
	"github.com/go-chi/jwtauth"
)

type AuthContext struct {
	UserID string
	Email  string
	Role   string
}
type contextKey string

const AuthKey contextKey = "auth"

func AuthMiddleware(jwtService *jwt.JwtService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			tokenString := jwtauth.TokenFromHeader(r)
			if tokenString == "" {
				http.Error(w, "missing token", http.StatusUnauthorized)
				return
			}

			claims, err := jwtService.Validate(tokenString)
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			auth := AuthContext{
				UserID: claims["user_id"].(string),
				Email:  claims["email"].(string),
				Role:   claims["role"].(string),
			}

			ctx := context.WithValue(r.Context(), AuthKey, auth)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

package middleware

import (
	Service "GCT/Structure/Services"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

var jwtSecret = []byte("super-secret-token-for-testing-jwt@@@")

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Service.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Store user info in request context
		ctx := context.WithValue(r.Context(), "USER", claims.Username)
		next(w, r.WithContext(ctx))
	}
}

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Extract token after "Bearer "
		tokenString := authHeader[len("Bearer "):]

		// Store token in request context
		ctx := context.WithValue(r.Context(), "USER_TOKEN", tokenString)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

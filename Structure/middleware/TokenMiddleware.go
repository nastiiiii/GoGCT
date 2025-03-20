package middleware

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

// for demo keep here
var jwtSecret = []byte("super-secret-token-for-testing-jwt@@@")
var jwtToken = []byte("super-secret-token-for-testing-jwt@@@")

// Claims for creating token based on the username
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken to create token and get all the information for it
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{Username: username, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtToken)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateJWT used with frontend, in order before doing the request the token will be taken from the header of the request
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

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

// AuthenticateMiddleware checks if the token valid
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

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtToken, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

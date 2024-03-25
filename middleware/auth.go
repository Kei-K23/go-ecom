package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Kei-K23/go-ecom/config"
	"github.com/Kei-K23/go-ecom/services/auth"
	"github.com/Kei-K23/go-ecom/utils"
	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const ClaimsContextKey ContextKey = "claims"

func CheckAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		// Check if Authorization header is present
		if authHeader == "" {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("authorization header is missing"))
			return
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		tokenString = strings.TrimSpace(tokenString)

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &auth.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Env.Secret), nil
		})

		// Check for token parsing errors
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("failed to parse token: %v", err))
			return
		}

		// Validate the token
		if !token.Valid {
			return
		}

		// Extract claims from the token
		claims, ok := token.Claims.(*auth.JWTClaim)
		if !ok {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token claims"))
			return
		}

		// Add the claims to the request context
		ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
		// Call the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

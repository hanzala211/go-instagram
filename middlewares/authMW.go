package middlewares

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hanzala211/instagram/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(utils.GetEnv("JWT_SECRET", "secret")), nil
			})
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		userId := claims["userId"].(string)
		ctx := context.WithValue(r.Context(), "userId", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
		})
}
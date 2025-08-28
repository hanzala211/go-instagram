package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hanzala211/instagram/internal/api/models"
	"github.com/hanzala211/instagram/internal/cache"
	"github.com/hanzala211/instagram/internal/services"
	"github.com/hanzala211/instagram/utils"
	"github.com/redis/go-redis/v9"
)
func AuthMiddleware(rdRepo *cache.RedisRepo, userService *services.UserService) func (next http.Handler) http.Handler{
	return func (next http.Handler) http.Handler {
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
		fmtStr := fmt.Sprintf("user-%s", userId)
		user, err := rdRepo.Get(fmtStr)
		fmt.Printf(user)
		var userData *models.User
		if err != nil  {
			if err == redis.Nil {
				userData, err = userService.GetUserById(userId)
				if err != nil {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
				err = rdRepo.Set(fmtStr, userData, time.Hour * 24)
				if err != nil {
					fmt.Printf("Error setting user data in redis: %v", err)
				}
			}else{
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}
		err = json.Unmarshal([]byte(user), &userData)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user", userData)
		next.ServeHTTP(w, r.WithContext(ctx))
		})
}
}
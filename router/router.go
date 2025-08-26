package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/hanzala211/instagram/internal/api/handler"
	"github.com/hanzala211/instagram/middlewares"
)

func SetupRouter(userHandler *handler.UserHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(u chi.Router) {
			u.Post("/signup", userHandler.Signup)
			u.Post("/login", userHandler.Login)	
			u.Group(func(r chi.Router) {
				r.Use(middlewares.AuthMiddleware)
				r.Get("/me", userHandler.ME)
			})
		})
	})
	return r
}
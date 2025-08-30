package router

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/hanzala211/instagram/internal/api/handler"
	"github.com/hanzala211/instagram/internal/cache"
	"github.com/hanzala211/instagram/internal/services"
	"github.com/hanzala211/instagram/middlewares"
)

func SetupRouter(userHandler *handler.UserHandler, postHandler *handler.PostHandler, rdRepo *cache.RedisRepo, userService *services.UserService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.tmpl")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := map[string]string{
			"Title":   "Home Page",
			"Message": "Go Instagram",
		}
		tmpl.Execute(w, data)
	})
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(u chi.Router) {
			u.Post("/signup", userHandler.Signup)
			u.Post("/login", userHandler.Login)
			u.Group(func(r chi.Router) {
				r.Use(middlewares.AuthMiddleware(rdRepo, userService))
				r.Get("/me", userHandler.ME)
			})
		})
		r.Route("/posts", func(u chi.Router) {
			u.Group(func(a chi.Router) {
				a.Use(middlewares.AuthMiddleware(rdRepo, userService))
				a.Post("/", postHandler.CreatePost)
			})

		})
	})
	return r
}

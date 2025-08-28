package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/hanzala211/instagram/db"
	"github.com/hanzala211/instagram/internal/api/handler"
	"github.com/hanzala211/instagram/internal/cache"
	"github.com/hanzala211/instagram/internal/repo"
	"github.com/hanzala211/instagram/internal/services"
	"github.com/hanzala211/instagram/router"
	"github.com/hanzala211/instagram/utils"
	"github.com/joho/godotenv"
)
	
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}
	
	database := db.ConnectPGDB()
	db.Migrations(database)
	
	log.Println("Initializing Redis connection...")
	rdRepo := cache.NewRedisClient()
	log.Println("Redis connection established successfully")
	
	
	userRepo := repo.NewUserRepo(database)
	store := repo.NewStorage(userRepo)
	userService := services.NewUserService(store)
	userHandler := handler.NewUserHandler(userService, rdRepo)
	router := router.SetupRouter(userHandler, rdRepo, userService)

	fmt.Println("Starting authentication service")
	port := utils.GetEnv("PORT", "4001")
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	err = http.ListenAndServe(port, router)
	if err != nil {
		fmt.Println("Error starting authentication service")
		panic(err)
	}
}	

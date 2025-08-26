package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hanzala211/instagram/db"
	"github.com/hanzala211/instagram/internal/api/handler"
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
	userRepo := repo.NewUserRepo(database)
	store := repo.NewStorage(userRepo)
	userService := services.NewUserService(store)
	userHandler := handler.NewUserHandler(userService)
	router := router.SetupRouter(userHandler)

	fmt.Println("Starting authentication service")
	err = http.ListenAndServe(utils.GetEnv("PORT", ":4001"), router)
	if err != nil {
		fmt.Println("Error starting authentication service")
		panic(err)
	}
}	
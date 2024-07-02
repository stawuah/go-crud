package main

import (
	"fmt"
	"log"

	"github/stawuah/go-crud/config"
	"github/stawuah/go-crud/internal/modules/users/queries"
	"github/stawuah/go-crud/internal/modules/users/repository"
	service "github/stawuah/go-crud/internal/modules/users/services"
	"github/stawuah/go-crud/pkg/router"
)

// Example usage
func main() {

	db, err := config.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Use db.GetDB() for operations
	fmt.Println("Database connection established successfully")

	userQueries := queries.NewUserQueries(db.GetDB())

	// Initialize repository layer
	userRepo := repository.NewUserRepository(userQueries)

	// Initialize service layer
	userService := service.NewUserService(userRepo)

	// Initialize and start HTTP server
	router := router.NewRouter(userService)
	log.Fatal(router.Run(":8080"))
}

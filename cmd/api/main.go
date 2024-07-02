package main

import (
	"fmt"
	"log"

	"github/stawuah/go-crud/config"
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
}

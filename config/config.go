package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sqlx.DB
}

// NewDatabase initializes a new database connection
func NewDatabase() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}

	connectionString := os.Getenv("DB_CONN_STRING")
	if connectionString == "" {
		return nil, fmt.Errorf("DB_CONN_STRING environment variable not set")
	}

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return &Database{db: db}, nil
}

// Close closes the database connection
func (d *Database) Close() error {
	return d.db.Close()
}

// GetDB returns the underlying sqlx.DB instance
func (d *Database) GetDB() *sqlx.DB {
	return d.db
}

package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/ruziba3vich/parking_searcher/pkg/config"
)

// ConnectDB initializes and returns a DB connection
func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %v", err)
		return nil, err
	}

	log.Println("Database connected successfully!")
	return db, nil
}

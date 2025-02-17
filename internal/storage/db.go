package storage

import (
	"database/sql"
	"log"

	"github.com/Masterminds/squirrel"
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

func NewStorage[T any](db *sql.DB) any {
	switch any(new(T)).(type) {
	case *CardStorage:
		return *&CardStorage{
			db: db,
			qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
		}
	case *HistoryStorage:
		return &HistoryStorage{
			db: db,
			qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
		}
	case *ParkStorage:
		return &ParkStorage{
			db: db,
			qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
		}
	case *SpotStorage:
		return &SpotStorage{
			db: db,
			qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
		}
	case *UserStorage:
		return &UserStorage{
			db: db,
			qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
		}
	default:
		return *new(T)
	}
}

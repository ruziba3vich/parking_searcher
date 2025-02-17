package storage

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/ruziba3vich/parking_searcher/internal/models"
)

type ParkStorage struct {
	db *sql.DB
	qb sq.StatementBuilderType
}

// NewParkStorage initializes ParkStorage with a DB connection
func NewParkStorage(db *sql.DB) *ParkStorage {
	return &ParkStorage{
		db: db,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

// Create a new park
func (s *ParkStorage) CreatePark(ctx context.Context, park models.Park) error {
	query, args, err := s.qb.Insert("parks").
		Columns("id", "name", "location", "capacity", "is_deleted").
		Values(park.ParkID, park.ParkName, park.Address, park.TotalSpotsCount, false).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	return err
}

// Get all parks (excluding deleted)
func (s *ParkStorage) GetAllParks(ctx context.Context) ([]models.Park, error) {
	query, args, err := s.qb.Select("id", "name", "location", "capacity").
		From("parks").
		Where(sq.Eq{"is_deleted": false}).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parks []models.Park
	for rows.Next() {
		var park models.Park
		if err := rows.Scan(&park.ParkID, &park.ParkName, &park.Address, &park.TotalSpotsCount); err != nil {
			return nil, err
		}
		parks = append(parks, park)
	}

	return parks, nil
}

// Find park by ID (excluding deleted)
func (s *ParkStorage) GetParkByID(ctx context.Context, id string) (*models.Park, error) {
	query, args, err := s.qb.Select("id", "name", "location", "capacity").
		From("parks").
		Where(sq.Eq{"id": id, "is_deleted": false}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var park models.Park
	err = s.db.QueryRowContext(ctx, query, args...).Scan(&park.ParkID, &park.ParkName, &park.Address, &park.TotalSpotsCount)
	if err != nil {
		return nil, err
	}

	return &park, nil
}

// Update park details
func (s *ParkStorage) UpdatePark(ctx context.Context, id string, data map[string]interface{}) error {
	query, args, err := s.qb.Update("parks").
		SetMap(data).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	return err
}

// Soft delete a park (set is_deleted = true)
func (s *ParkStorage) DeletePark(ctx context.Context, id string) error {
	query, args, err := s.qb.Update("parks").
		Set("is_deleted", true).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	return err
}

/*
	type Park struct {
		ParkID                   uuid.UUID `json:"park_id"`
		ParkName                 string    `json:"park_name"`
		Address                  string    `json:"address"`
		PricePerHour             float64   `json:"price_ph"`
		Status                   string    `json:"status"`
		AvailableSpotsCount      int       `json:"available_spots_count"`
		TotalSpotsCount          int       `json:"total_spots_count"`
		ElectroChargingAvailable bool      `json:"electro_charging_available"`
		Rating                   float64   `json:"rating"`
		ParkBalance              float64   `json:"park_balance"`
		Latitude                 float64   `json:"latitude"`
		Longitude                float64   `json:"longitude"`
	}
*/

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
		Columns(
			"park_name",
			"park_name",
			"address",
			"price_ph",
			"status",
			"available_spots_count",
			"total_spots_count",
			"electro_charging_available",
			"latitude",
			"longitude",
			"is_deleted").
		Values(
			park.ParkID,
			park.ParkName,
			park.Address,
			park.PricePerHour,
			park.Status,
			park.AvailableSpotsCount,
			park.TotalSpotsCount,
			park.ElectroChargingAvailable,
			park.Longitude,
			park.Latitude,
			park.TotalSpotsCount, false).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	return err
}

// GetAllParks retrieves all parks (excluding deleted) with pagination.
func (s *ParkStorage) GetAllParks(ctx context.Context, limit, offset int) ([]models.Park, error) {
	query, args, err := s.qb.Select(
		"park_id",
		"park_name",
		"address",
		"price_ph",
		"status",
		"available_spots_count",
		"total_spots_count",
		"electro_charging_available",
		"rating",
		"latitude",
		"longitude",
	).
		From("parks").
		Where(sq.Eq{"is_deleted": false}).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
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
		if err := rows.Scan(
			&park.ParkID,
			&park.ParkName,
			&park.Address,
			&park.PricePerHour,
			&park.Status,
			&park.AvailableSpotsCount,
			&park.TotalSpotsCount,
			&park.ElectroChargingAvailable,
			&park.Rating,
			&park.Latitude,
			&park.Longitude,
		); err != nil {
			return nil, err
		}
		parks = append(parks, park)
	}

	return parks, nil
}

// Find park by ID (excluding deleted)
func (s *ParkStorage) GetParkByID(ctx context.Context, id string) (*models.Park, error) {
	query, args, err := s.qb.Select(
		"park_id",
		"park_name",
		"address",
		"price_ph",
		"status",
		"available_spots_count",
		"total_spots_count",
		"electro_charging_available",
		"rating",
		"latitude",
		"longitude",
	).
		From("parks").
		Where(sq.Eq{"id": id, "is_deleted": false}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var park models.Park
	err = s.db.QueryRowContext(ctx, query, args...).Scan(
		&park.ParkID,
		&park.ParkName,
		&park.Address,
		&park.PricePerHour,
		&park.Status,
		&park.AvailableSpotsCount,
		&park.TotalSpotsCount,
		&park.ElectroChargingAvailable,
		&park.Rating,
		&park.Latitude,
		&park.Longitude,
	)
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
	park_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    park_name VARCHAR(64) NOT NULL,
    address VARCHAR(166) NOT NULL,
    price_ph DOUBLE PRECISION NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('available', 'full', 'closed')),
    available_spots_count INT NOT NULL,
    total_spots_count INT NOT NULL,
    electro_charging_available BOOLEAN NOT NULL DEFAULT FALSE,
    rating DOUBLE PRECISION DEFAULT 0,
    park_balance DOUBLE PRECISION DEFAULT 0,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE
*/

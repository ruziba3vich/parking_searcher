package storage

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ruziba3vich/parking_searcher/internal/models"
)

type SpotStorage struct {
	db *sql.DB
	qb sq.StatementBuilderType
}

func (s *SpotStorage) CreateSpot(ctx context.Context, spot *models.Spot) (uuid.UUID, error) {
	query, args, err := s.qb.Insert("spots").
		Columns(
			"is_available",
			"electro_charger_available",
			"vehicle_type",
			"booked_from",
			"booked_till",
			"about",
		).
		Values(
			spot.IsAvailable,
			spot.ElectroChargerAvailable,
			spot.VehicleType,
			spot.BookedFrom,
			spot.BookedTill,
			spot.About,
		).
		ToSql()
	if err != nil {
		return uuid.Nil, err
	}

	query += " RETURNING spot_id;"

	// Execute the insert query and retrieve the SpotID
	var spotID uuid.UUID
	err = s.db.QueryRowContext(ctx, query, args...).Scan(&spotID)
	if err != nil {
		return uuid.Nil, err
	}

	return spotID, nil
}

func (s *SpotStorage) GetSpot(ctx context.Context, spotID string) (*models.Spot, error) {
	query, args, err := s.qb.Select(
		"spot_id",
		"park_id",
		"is_available",
		"electro_charger_available",
		"vehicle_type",
		"booked_from",
		"booked_till",
		"about",
	).
		From("spots").
		Where(sq.Eq{"spot_id": spotID}, sq.Eq{"is_deleted": false}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)

	var spot models.Spot
	err = row.Scan(
		&spot.SpotID,
		&spot.ParkID,
		&spot.IsAvailable,
		&spot.ElectroChargerAvailable,
		&spot.VehicleType,
		&spot.BookedFrom,
		&spot.BookedTill,
		&spot.About)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("spot not found")
		}
		return nil, err
	}

	return &spot, nil
}

func (s *SpotStorage) UpdateSpot(ctx context.Context, spotID string, data map[string]interface{}) error {
	// Make sure that "is_deleted" is not included in the update
	if _, ok := data["is_deleted"]; ok {
		delete(data, "is_deleted")
	}

	query, args, err := s.qb.Update("spots").
		SetMap(data).
		Where(sq.Eq{"spot_id": spotID}, sq.Eq{"is_deleted": false}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	return err
}

func (s *SpotStorage) GetAllSpotsByParkId(ctx context.Context, parkID string, limit int, offset int) ([]*models.Spot, error) {
	query, args, err := s.qb.Select(
		"spot_id",
		"park_id",
		"is_available",
		"electro_charger_available",
		"vehicle_type",
		"booked_from",
		"booked_till",
		"about",
	).
		From("spots").
		Where(sq.Eq{"park_id": parkID}, sq.Eq{"is_deleted": false}).
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

	var spots []*models.Spot
	for rows.Next() {
		var spot models.Spot
		if err := rows.Scan(
			&spot.SpotID,
			&spot.ParkID,
			&spot.IsAvailable,
			&spot.ElectroChargerAvailable,
			&spot.VehicleType,
			&spot.BookedFrom,
			&spot.BookedTill,
			&spot.About); err != nil {
			return nil, err
		}
		spots = append(spots, &spot)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return spots, nil
}

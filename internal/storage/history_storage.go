package storage

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ruziba3vich/parking_searcher/internal/models"
)

type (
	HistoryStorage struct {
		db *sql.DB
		qb sq.StatementBuilderType
	}
)

func (s *HistoryStorage) CreateHistory(ctx context.Context, history *models.History) (uuid.UUID, error) {
	// Use squirrel to build the insert query
	query, args, err := s.qb.Insert("history").
		Columns(
			"user_id",
			"park_id",
			"spot_id",
			"start_time",
			"end_time",
			"total_price",
			"payment_type",
			"rate",
			"comment",
		).
		Values(
			history.HistoryID,
			history.UserID,
			history.ParkID,
			history.SpotId,
			history.StartTime,
			history.EndTime,
			history.TotalPrice,
			history.PaymentType,
			history.Rate,
			history.Comment,
		).
		ToSql()
	if err != nil {
		return uuid.Nil, err
	}

	// Execute the query
	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return uuid.Nil, err
	}

	return history.HistoryID, nil
}

func (s *HistoryStorage) GetAllHistory(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*models.History, error) {
	// Build the query based on the provided filter
	queryBuilder := s.qb.Select(
		"history_id",
		"user_id",
		"park_id",
		"spot_id",
		"start_time",
		"end_time",
		"total_price",
		"payment_type",
		"rate",
		"comment",
	).
		From("history").
		Where(sq.Eq{"is_deleted": false})

	// Dynamically add the filter conditions
	for key, value := range filter {
		queryBuilder = queryBuilder.Where(sq.Eq{key: value})
	}

	// Apply pagination
	queryBuilder = queryBuilder.Limit(uint64(limit)).Offset(uint64(offset))

	// Build the final SQL query
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	// Execute the query
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []*models.History
	for rows.Next() {
		var history models.History
		if err := rows.Scan(
			&history.HistoryID,
			&history.UserID,
			&history.ParkID,
			&history.SpotId,
			&history.StartTime,
			&history.EndTime,
			&history.TotalPrice,
			&history.PaymentType,
			&history.Rate,
			&history.Comment); err != nil {
			return nil, err
		}
		histories = append(histories, &history)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return histories, nil
}

func (s *HistoryStorage) DeleteHistory(ctx context.Context, historyID string) error {
	// Use squirrel to build the update query
	query, args, err := s.qb.Update("history").
		Set("is_deleted", true).
		Where(sq.Eq{"history_id": historyID}).
		ToSql()
	if err != nil {
		return err
	}

	// Execute the query
	_, err = s.db.ExecContext(ctx, query, args...)
	return err
}

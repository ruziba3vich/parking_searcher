package storage

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ruziba3vich/parking_searcher/internal/models"
)

type (
	CardStorage struct {
		db *sql.DB
		qb sq.StatementBuilderType
	}
)

func (s *CardStorage) CreateCard(ctx context.Context, card *models.Card) (uuid.UUID, error) {
	// Use squirrel to build the insert query
	query, args, err := s.qb.Insert("cards").
		Columns("user_id", "card_number", "balance").
		Values(card.CardID, card.UserID, card.CardNumber, card.Balance).
		ToSql()
	if err != nil {
		return uuid.Nil, err
	}

	query += " RETURNING card_id;"
	// Execute the query
	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return uuid.Nil, err
	}

	return card.CardID, nil
}

func (s *CardStorage) DeleteCard(ctx context.Context, cardID string) error {
	// Use squirrel to build the update query for soft delete
	query, args, err := s.qb.Update("cards").
		Set("is_deleted", true).
		Where(sq.Eq{"card_id": cardID}).
		ToSql()
	if err != nil {
		return err
	}

	// Execute the query
	_, err = s.db.ExecContext(ctx, query, args...)
	return err
}

func (s *CardStorage) GetCardsByUserId(ctx context.Context, userID string) ([]*models.Card, error) {
	// Build the query to fetch all cards for the given user_id
	query, args, err := s.qb.Select("card_id", "user_id", "card_number", "balance").
		From("cards").
		Where(sq.Eq{"user_id": userID, "is_deleted": false}).
		ToSql()
	if err != nil {
		return nil, err
	}

	// Execute the query
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []*models.Card
	for rows.Next() {
		var card models.Card
		if err := rows.Scan(&card.CardID, &card.UserID, &card.CardNumber, &card.Balance); err != nil {
			return nil, err
		}
		cards = append(cards, &card)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

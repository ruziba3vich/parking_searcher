package storage

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ruziba3vich/parking_searcher/internal/models"
)

type (
	UserStorage struct {
		db *sql.DB
		qb sq.StatementBuilderType
	}
)

func (s *UserStorage) CreateUser(ctx context.Context, user models.User) (uuid.UUID, error) {
	// Use squirrel to build the insert query, excluding the user_id
	query, args, err := s.qb.Insert("users").
		Columns("email", "password", "full_name", "phone").
		Values(user.Email, user.Password, user.FullName, user.Phone).
		ToSql()
	if err != nil {
		return uuid.Nil, err
	}

	// Modify the query to use RETURNING to retrieve the generated user_id
	query = query + " RETURNING user_id" // Add RETURNING clause to get the user_id

	// Execute the query and retrieve the generated user_id
	var userID uuid.UUID
	err = s.db.QueryRowContext(ctx, query, args...).Scan(&userID)
	if err != nil {
		return uuid.Nil, err
	}

	// Set the generated user_id back to the user object
	user.UserID = userID

	return userID, nil
}

func (s *UserStorage) UpdateUser(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error {
	// Define valid fields for a User object
	validFields := map[string]bool{
		"email":     true,
		"full_name": true,
		"phone":     true,
		// "password" is excluded as it's often sensitive
	}

	// Filter out invalid fields from the data map
	filteredData := make(map[string]interface{})
	for key, value := range data {
		if validFields[key] {
			filteredData[key] = value
		} else {
			continue
		}
	}

	// Ensure there are fields to update
	if len(filteredData) == 0 {
		return fmt.Errorf("no valid fields provided for update")
	}

	// Build the query with squirrel for the update
	query, args, err := s.qb.Update("users").
		SetMap(filteredData).
		Where(sq.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return err
	}

	// Execute the update query
	_, err = s.db.ExecContext(ctx, query, args...)
	return err
}

func (s *UserStorage) GetUserById(ctx context.Context, userID string) (*models.User, error) {
	// Use squirrel to build the select query
	query, args, err := s.qb.Select("user_id", "email", "full_name", "phone").
		From("users").
		Where(sq.Eq{"user_id": userID, "is_deleted": false}).
		ToSql()
	if err != nil {
		return nil, err
	}

	// Execute the query
	row := s.db.QueryRowContext(ctx, query, args...)

	var user models.User
	err = row.Scan(&user.UserID, &user.Email, &user.FullName, &user.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (s *UserStorage) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	// Build the update query to set the is_deleted field to true
	query, args, err := s.qb.Update("users").
		Set("is_deleted", true).
		Where(sq.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return err
	}

	// Execute the update query
	_, err = s.db.ExecContext(ctx, query, args...)
	return err
}

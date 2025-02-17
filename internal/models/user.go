package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	FullName string    `json:"full_name"`
	Phone    string    `json:"phone"`
}

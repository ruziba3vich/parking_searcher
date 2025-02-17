package models

import (
	"github.com/google/uuid"
)

type Card struct {
	CardID     uuid.UUID `json:"card_id" gorm:"type:uuid;primaryKey"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;index"`
	CardNumber string    `json:"card_number"`
	Balance    float64   `json:"balance"`
}

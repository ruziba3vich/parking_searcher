package models

import (
	"time"

	"github.com/google/uuid"
)

type History struct {
	HistoryID   uuid.UUID  `json:"history_id"`
	UserID      uuid.UUID  `json:"user_id"`
	ParkID      uuid.UUID  `json:"park_id"`
	SpotId      uuid.UUID  `json:"spot_id"`
	StartTime   time.Time  `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	TotalPrice  float64    `json:"total_price"`
	PaymentType string     `json:"payment_type"`
	Rate        *float32   `json:"rate"`
	Comment     *string    `json:"comment"`
}

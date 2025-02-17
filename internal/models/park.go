package models

import (
	"github.com/google/uuid"
)

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

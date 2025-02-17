package models

import (
	"time"

	"github.com/google/uuid"
)

type Spot struct {
	SpotID                  uuid.UUID   `json:"spot_id" gorm:"type:uuid;primaryKey"`
	ParkID                  uuid.UUID   `json:"park_id" gorm:"type:uuid;index"`
	IsAvailable             bool        `json:"is_available" gorm:"default:true"`
	ElectroChargerAvailable bool        `json:"electro_charger_available"`
	VehicleType             VehicleType `json:"vehicle_type"`
	BookedFrom              *time.Time  `json:"booked_from"`
	BookedTill              *time.Time  `json:"booked_till"`
	About                   string      `json:"about"`
}

type VehicleType string

const (
	CAR   VehicleType = "car"
	TRUCK VehicleType = "truck"
)

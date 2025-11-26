package entities

import (
	"time"
)

type Booking struct {
	ID uint
	Date time.Time
	StartAt time.Time
	EndAt time.Time
}
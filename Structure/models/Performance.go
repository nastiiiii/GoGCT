package models

import (
	"time"
)

type Performance struct {
	PerformanceId          int
	SeatBandPricing        string
	SeatAvailability       string
	PerformanceName        string
	PerformanceDescription string
	PerformanceDate        time.Time
	PerformanceActors      string
}

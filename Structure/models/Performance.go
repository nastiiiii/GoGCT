package models

import (
	"time"
)

// Performance represents entity from the database
type Performance struct {
	PerformanceId          int
	SeatBandPricing        string
	SeatAvailability       string // have all seats with their availability
	PerformanceName        string
	PerformanceDescription string
	PerformanceDate        time.Time
	PerformanceActors      string
}

func NewPerformance(seatBandPricing, seatAvailability, performanceName, performanceDescription, performanceActors string, performanceDate time.Time) Performance {
	return Performance{
		SeatBandPricing:        seatBandPricing,
		SeatAvailability:       seatAvailability,
		PerformanceName:        performanceName,
		PerformanceDescription: performanceDescription,
		PerformanceDate:        performanceDate,
		PerformanceActors:      performanceActors,
	}
}

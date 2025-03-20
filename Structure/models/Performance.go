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

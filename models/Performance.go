package models

import "time"

type Performance struct {
	//performanceId          int
	seatBandPricing        string
	seatAvailability       string
	performanceName        string
	performanceDescription string
	performanceDate        time.Time
	performanceActors      string
}

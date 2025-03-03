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

type IPerformanceService interface {
	updatePerformance(performance Performance, performanceId int)
	getPerformance(performanceId int) Performance
	getPerformanceByName(performanceName string) Performance
	getPerformances() []Performance
}

package Interfaces

import Models "GCT/Structure/models"

type IPerformanceService interface {
	CreatePerformance(p Models.Performance) (int, error)
	DeletePerformance(id int) bool
	UpdatePerformance(performance Models.Performance, performanceId int) (Models.Performance, error)
	GetPerformanceById(performanceId int) (Models.Performance, error)
	GetPerformanceByName(performanceName string) (Models.Performance, error)
	GetPerformances() []Models.Performance
	GetAllSeats(performance Models.Performance) (map[string][]Models.Seats, error)
	GetAvailableSeats(performance Models.Performance) (map[string][]Models.Seats, error)
	GetSeatPrice(performance Models.Performance, seatBand string) (float64, error)
	GetPerformancePrice(performance Models.Performance) (map[string]float64, error)
	ChangeSeatAvailability(performance *Models.Performance, seatBand string, seatNumber int, status bool) (bool, error)
}

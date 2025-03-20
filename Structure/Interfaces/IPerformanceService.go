package Interfaces

import Models "GCT/Structure/models"

// IPerformanceService Description: Shows all methods which PerformanceService is implemented and to which the controllers have access through dependency injection
type IPerformanceService interface {
	//Create
	CreatePerformance(p Models.Performance) (int, error)
	//Delte
	DeletePerformance(id int) bool
	//Update
	UpdatePerformance(performance Models.Performance, performanceId int) (Models.Performance, error)
	//Get
	GetPerformanceById(performanceId int) (Models.Performance, error)
	GetPerformanceByName(performanceName string) (Models.Performance, error)
	GetPerformances() []Models.Performance
	GetAllSeats(performance Models.Performance) (map[string][]Models.Seats, error)
	GetAvailableSeats(performance Models.Performance) (map[string][]Models.Seats, error)
	GetSeatPrice(performance Models.Performance, seatBand string) (float64, error)
	GetPerformancePrice(performance Models.Performance) (map[string]float64, error)
	// Update related to seats
	ChangeSeatAvailability(performance *Models.Performance, seatBand string, seatNumber int, status bool) (bool, error)
}

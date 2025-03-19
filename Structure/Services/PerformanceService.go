package Services

import (
	"GCT/Structure/Util"
	Models "GCT/Structure/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"strings"
)

type PerformanceService struct {
	DB *pgx.Conn
}

// Approved
func (s PerformanceService) UpdatePerformance(performance Models.Performance, performanceId int) (Models.Performance, error) {
	query := `UPDATE "Performances" 
		SET "seatBandPricing" = $1, 
		    "seatAvailibility" = $2, 
		    "performanceName" = $3, 
		    "performanceDescription" = $4, 
		    "performanceDate" = $5, 
		    "perofrmanceActors" = $6 
		WHERE "performanceID" = $7`

	_, err := s.DB.Exec(
		context.Background(), query,
		&performance.SeatBandPricing,
		&performance.SeatAvailability,
		&performance.PerformanceName,
		&performance.PerformanceDescription,
		&performance.PerformanceDate,
		&performance.PerformanceActors,
		performanceId)

	updatedPerformance, err := s.GetPerformanceById(performanceId)

	if err != nil {
		log.Println("Error updating performance:", err)
		return Models.Performance{}, err
	}
	return updatedPerformance, nil
}

// Approved
func (s PerformanceService) GetPerformanceById(performanceId int) (Models.Performance, error) {
	var performance Models.Performance
	query := `SELECT "performanceID" ,"seatBandPricing", "seatAvailibility", "performanceName", 
	                 "performanceDescription", "performanceDate", "perofrmanceActors"  FROM "Performances" WHERE "performanceID" = $1`

	err := s.DB.QueryRow(context.Background(), query, performanceId).Scan(
		&performance.PerformanceId,
		&performance.SeatBandPricing,
		&performance.SeatAvailability,
		&performance.PerformanceName,
		&performance.PerformanceDescription,
		&performance.PerformanceDate,
		&performance.PerformanceActors,
	)
	if err != nil {
		return Models.Performance{}, errors.New("problem in finding performance")
	}
	return performance, nil
}

// Approved
func (s PerformanceService) GetPerformanceByName(performanceName string) (Models.Performance, error) {
	var performance Models.Performance
	query := `SELECT "performanceID" ,"seatBandPricing", "seatAvailibility", "performanceName", 
	                 "performanceDescription", "performanceDate", "perofrmanceActors"  FROM "Performances" WHERE "performanceName" = $1`

	err := s.DB.QueryRow(context.Background(), query, performanceName).Scan(
		&performance.PerformanceId,
		&performance.SeatBandPricing,
		&performance.SeatAvailability,
		&performance.PerformanceName,
		&performance.PerformanceDescription,
		&performance.PerformanceDate,
		&performance.PerformanceActors,
	)
	if err != nil {
		return Models.Performance{}, errors.New("problem in finding performance")
	}
	return performance, nil
}

// Approved
func (s PerformanceService) GetPerformances() []Models.Performance {
	var performances []Models.Performance
	query := `SELECT "performanceID" ,"seatBandPricing", "seatAvailibility", "performanceName", 
	                 "performanceDescription", "performanceDate", "perofrmanceActors"  FROM "Performances" ORDER BY $1`
	rows, err := s.DB.Query(context.Background(), query, "performanceDate")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var performance Models.Performance
		err = rows.Scan(
			&performance.PerformanceId,
			&performance.SeatBandPricing,
			&performance.SeatAvailability,
			&performance.PerformanceName,
			&performance.PerformanceDescription,
			&performance.PerformanceDate,
			&performance.PerformanceActors)
		if err != nil {
			log.Fatal(err)
		}
		performances = append(performances, performance)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return performances
}

// Approved
func (s PerformanceService) CreatePerformance(p Models.Performance) (int, error) {
	query := `INSERT INTO "Performances" 
		("seatBandPricing", "seatAvailibility", "performanceName", 
		 "performanceDescription", "performanceDate", "perofrmanceActors") 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING "performanceID"`

	var id int
	err := s.DB.QueryRow(
		context.Background(), query,
		p.SeatBandPricing, p.SeatAvailability,
		p.PerformanceName, p.PerformanceDescription,
		p.PerformanceDate, p.PerformanceActors,
	).Scan(&id)

	if err != nil {
		log.Println("Error inserting performance:", err)
		return 0, err
	}
	return id, nil
}

// Approved
func (s PerformanceService) DeletePerformance(id int) bool {
	query := `DELETE FROM "Performances" WHERE "performanceID" = $1`

	_, err := s.DB.Exec(context.Background(), query, id)
	if err != nil {
		log.Println("Error deleting performance:", err)
		return false
	}
	return true
}

// Approved
func (s PerformanceService) GetAllSeats(performance Models.Performance) (map[string][]Models.Seats, error) {
	return Util.ParseSeatAvailability(performance.SeatAvailability)
}

// Approved
func (s PerformanceService) GetAvailableSeats(performance Models.Performance) (map[string][]Models.Seats, error) {
	seatMap, err := Util.ParseSeatAvailability(performance.SeatAvailability)
	if err != nil {
		return nil, err
	}

	// Filter only available seats
	availableSeats := make(map[string][]Models.Seats)
	for band, seats := range seatMap {
		for _, seat := range seats {
			if seat.Availability {
				availableSeats[band] = append(availableSeats[band], seat)
			}
		}
	}

	return availableSeats, nil
}

// Approved
func (s PerformanceService) GetSeatPrice(performance Models.Performance, seatBand string) (float64, error) {
	prices, err := Util.ParseSeatPrices(performance.SeatBandPricing)
	if err != nil {
		return 0, err
	}

	price, exists := prices[seatBand]
	if !exists {
		return 0, errors.New("seat band not found")
	}

	return price, nil
}

// Approved
func (s PerformanceService) GetPerformancePrice(performance Models.Performance) (map[string]float64, error) {
	return Util.ParseSeatPrices(performance.SeatBandPricing)
}

// Approved
func (s PerformanceService) ChangeSeatAvailability(performance *Models.Performance, seatBand string, seatNumber int, status bool) (bool, error) {
	seatMap, err := Util.ParseSeatAvailability(performance.SeatAvailability)
	if err != nil {
		return false, err
	}

	seats, exists := seatMap[seatBand]
	if !exists {
		return false, errors.New("seat band not found")
	}

	// Update seat availability
	for i, seat := range seats {
		if seat.Seat == seatNumber {
			seatMap[seatBand][i].Availability = status
			break
		}
	}

	// Convert back to string format for saving
	var updatedSeats []string
	for band, seats := range seatMap {
		var seatEntries []string
		for _, seat := range seats {
			seatEntries = append(seatEntries, fmt.Sprintf("%d-%t", seat.Seat, seat.Availability))
		}
		updatedSeats = append(updatedSeats, fmt.Sprintf("%s: %s", band, strings.Join(seatEntries, ", ")))
	}

	// Update the performance struct
	performance.SeatAvailability = strings.Join(updatedSeats, "; ")

	_, err = s.UpdatePerformance(*performance, performance.PerformanceId)
	if err != nil {
		log.Println("Error updating seat availability in database:", err)
	}

	return true, nil
}

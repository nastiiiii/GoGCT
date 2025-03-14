package Services

import (
	Models "GCT/Structure/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"strconv"
	"strings"
)

type Seats struct {
	Seat         int
	Availability bool
}

type PerformanceService struct {
	DB *pgx.Conn
}

type IPerformanceService interface {
	CreatePerformance(p Models.Performance) (int, error)
	DeletePerformance(id int) bool
	UpdatePerformance(performance Models.Performance, performanceId int) (Models.Performance, error)
	GetPerformanceById(performanceId int) (Models.Performance, error)
	GetPerformanceByName(performanceName string) (Models.Performance, error)
	GetPerformances() []Models.Performance
	//five more methods
	GetAllSeats(performance Models.Performance) (map[string][]Seats, error)
	GetAvailableSeats(performance Models.Performance) (map[string][]Seats, error)
	GetSeatPrice(performance Models.Performance, seatBand string) (float64, error)
	GetPerformancePrice(performance Models.Performance) (map[string]float64, error)
	ChangeSeatAvailability(performance *Models.Performance, seatBand string, seatNumber int, status bool) (bool, error)
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
func (s PerformanceService) GetAllSeats(performance Models.Performance) (map[string][]Seats, error) {
	return ParseSeatAvailability(performance.SeatAvailability)
}

// Approved
func (s PerformanceService) GetAvailableSeats(performance Models.Performance) (map[string][]Seats, error) {
	seatMap, err := ParseSeatAvailability(performance.SeatAvailability)
	if err != nil {
		return nil, err
	}

	// Filter only available seats
	availableSeats := make(map[string][]Seats)
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
	prices, err := ParseSeatPrices(performance.SeatBandPricing)
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
	return ParseSeatPrices(performance.SeatBandPricing)
}

// Approved
func (s PerformanceService) ChangeSeatAvailability(performance *Models.Performance, seatBand string, seatNumber int, status bool) (bool, error) {
	seatMap, err := ParseSeatAvailability(performance.SeatAvailability)
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

// utils
// Parse seat band prices from "A:100; B:200; C:300;" format
func ParseSeatPrices(data string) (map[string]float64, error) {
	prices := make(map[string]float64)
	parts := strings.Split(data, ";") // ["A:100", " B:200", " C:300"]

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		bandData := strings.Split(part, ":") // ["A", "100"]
		if len(bandData) != 2 {
			return nil, errors.New("invalid seat price format")
		}

		price, err := strconv.ParseFloat(strings.TrimSpace(bandData[1]), 64)
		if err != nil {
			return nil, err
		}

		prices[strings.TrimSpace(bandData[0])] = price
	}

	return prices, nil
}

// Parse seat availability from "A: 12-true, 13-false, 14-true; B: 15-true, 16-true, 17-true;"
func ParseSeatAvailability(data string) (map[string][]Seats, error) {
	seatMap := make(map[string][]Seats)
	sections := strings.Split(data, ";") // ["A: 12-true, 13-false, 14-true", " B: 15-true, 16-true, 17-true"]

	for _, section := range sections {
		section = strings.TrimSpace(section)
		if section == "" {
			continue
		}

		parts := strings.Split(section, ":") // ["A", " 12-true, 13-false, 14-true"]
		if len(parts) != 2 {
			return nil, errors.New("invalid seat availability format")
		}

		seatBand := strings.TrimSpace(parts[0])
		seats := strings.Split(parts[1], ",") // [" 12-true", " 13-false", " 14-true"]

		for _, seat := range seats {
			seat = strings.TrimSpace(seat)
			seatParts := strings.Split(seat, "-") // ["12", "true"]
			if len(seatParts) != 2 {
				return nil, errors.New("invalid seat entry format")
			}

			seatNumber, err := strconv.Atoi(strings.TrimSpace(seatParts[0]))
			if err != nil {
				return nil, err
			}

			availability, err := strconv.ParseBool(strings.TrimSpace(seatParts[1]))
			if err != nil {
				return nil, err
			}

			seatMap[seatBand] = append(seatMap[seatBand], Seats{
				Seat:         seatNumber,
				Availability: availability,
			})
		}
	}

	return seatMap, nil
}

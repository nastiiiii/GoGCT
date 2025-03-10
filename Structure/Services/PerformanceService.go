package Services

import (
	Models "GCT/Structure/models"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

type PerformanceService struct {
	DB *pgx.Conn
}

type IPerformanceService interface {
	CreatePerformance(p Models.Performance) (int, error)
	DeletePerformance(id int) bool
	UpdatePerformance(performance Models.Performance, performanceId int) (int, error)
	GetPerformanceById(performanceId int) Models.Performance
	GetPerformanceByName(performanceName string) Models.Performance
	GetPerformances() []Models.Performance
}

func (s PerformanceService) UpdatePerformance(performance Models.Performance, performanceId int) (int, error) {
	query := `UPDATE performance 
		SET seat_band_pricing = $1, 
		    seat_availability = $2, 
		    performance_name = $3, 
		    performance_description = $4, 
		    performance_date = $5, 
		    performance_actors = $6 
		WHERE performance_id = $7`

	_, err := s.DB.Exec(
		context.Background(), query,
		performance.SeatBandPricing,
		performance.SeatAvailability,
		performance.PerformanceName,
		performance.PerformanceDescription,
		performance.PerformanceDate,
		performance.PerformanceActors,
		performanceId)

	if err != nil {
		log.Println("Error updating performance:", err)
		return 0, err
	}
	return performanceId, nil

}

func (s PerformanceService) GetPerformanceById(performanceId int) Models.Performance {
	var performance Models.Performance
	query := `SELECT * FROM performance WHERE performance_id = $1`

	err := s.DB.QueryRow(context.Background(), query, performanceId).Scan(&performance)
	if err != nil {
		log.Fatal(err)
	}
	return performance
}

func (s PerformanceService) GetPerformanceByName(performanceName string) Models.Performance {
	var performance Models.Performance
	query := `SELECT * FROM performance WHERE performance_name = $1`

	err := s.DB.QueryRow(context.Background(), query, performanceName).Scan(&performance)
	if err != nil {
		log.Fatal(err)
	}
	return performance
}

// TODO ORDER BY CHECK
func (s PerformanceService) GetPerformances() []Models.Performance {
	var performances []Models.Performance
	query := `SELECT * FROM performance ORDER BY $1`
	rows, err := s.DB.Query(context.Background(), query, "PerformanceDate")
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

func (s PerformanceService) CreatePerformance(p Models.Performance) (int, error) {
	query := `INSERT INTO performance 
		(seat_band_pricing, seat_availability, performance_name, 
		 performance_description, performance_date, performance_actors) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING performance_id`

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

func (s PerformanceService) DeletePerformance(id int) bool {
	query := `DELETE FROM performance WHERE performance_id = $1`

	_, err := s.DB.Exec(context.Background(), query, id)
	if err != nil {
		log.Println("Error deleting performance:", err)
		return false
	}
	return true
}

package models

import (
	"github.com/jackc/pgx/v5"
	"time"
)

type Performance struct {
	performanceId          int
	seatBandPricing        string
	seatAvailability       string
	performanceName        string
	performanceDescription string
	performanceDate        time.Time
	performanceActors      string
	service                IPerformanceService
}

type PerformanceService struct {
	DB *pgx.Conn
}

type IPerformanceService interface {
	CreatePerformance(p Performance) Performance
	CreatePerformanceByParams(seatBandPricing string, seatAvailability string, performanceName string, price string) *Performance
	DeletePerformance(id int) bool
	UpdatePerformance(performance Performance, performanceId int)
	GetPerformanceById(performanceId int) Performance
	GetPerformanceByName(performanceName string) Performance
	GetPerformances() []Performance
}

func (s PerformanceService) UpdatePerformance(performance Performance, performanceId int) {
	//TODO implement me
	panic("implement me")
}

func (s PerformanceService) GetPerformanceById(performanceId int) Performance {
	//TODO implement me
	panic("implement me")
}

func (s PerformanceService) GetPerformanceByName(performanceName string) Performance {
	//TODO implement me
	panic("implement me")
}

func (s PerformanceService) GetPerformances() []Performance {
	//TODO implement me
	panic("implement me")
}

func (s PerformanceService) CreatePerformance(p Performance) {
	//TODO implement me
	panic("implement me")
}

func (s PerformanceService) DeletePerformance(id int) bool {
	//TODO implement me
	panic("implement me")
}

func (s PerformanceService) CreatePerformanceByParams(seatBandPricing string, seatAvailability string, performanceName string, price string) *Performance {
	//TODO implement me
	panic("implement me")
}
